package globaltokenbucket

import (
	"sync"
	"time"

	"github.com/buraksezer/olric"
	"github.com/buraksezer/olric/config"

	distcache "github.com/fluxninja/aperture/v2/pkg/dist-cache"
	"github.com/fluxninja/aperture/v2/pkg/log"
	ratelimiter "github.com/fluxninja/aperture/v2/pkg/rate-limiter"
	"github.com/fluxninja/aperture/v2/pkg/utils"
)

const (
	// TakeNFunction is the name of the function used to take N tokens from the bucket.
	TakeNFunction = "TakeN"
)

// GlobalTokenBucket implements Limiter.
type GlobalTokenBucket struct {
	mu             sync.RWMutex
	dMap           *olric.DMap
	name           string
	bucketCapacity float64
	fillAmount     float64
	interval       time.Duration
	continuousFill bool
	passThrough    bool
	dc             *distcache.DistCache
}

// NewGlobalTokenBucket creates a new instance of DistCacheRateTracker.
func NewGlobalTokenBucket(dc *distcache.DistCache,
	name string,
	interval time.Duration,
	maxIdleDuration time.Duration,
	continuousFill bool,
) (*GlobalTokenBucket, error) {
	gtb := &GlobalTokenBucket{
		name:           name,
		interval:       interval,
		passThrough:    true,
		continuousFill: continuousFill,
		dc:             dc,
	}

	dmapConfig := config.DMap{
		MaxIdleDuration: maxIdleDuration,
		Functions: map[string]config.Function{
			TakeNFunction: gtb.takeN,
		},
	}

	dMap, err := dc.NewDMap(name, dmapConfig)
	if err != nil {
		return nil, err
	}

	gtb.dMap = dMap

	return gtb, nil
}

// SetBucketCapacity sets the rate limit for the rate limiter.
func (gtb *GlobalTokenBucket) SetBucketCapacity(bucketCapacity float64) {
	gtb.mu.Lock()
	defer gtb.mu.Unlock()
	gtb.bucketCapacity = bucketCapacity
}

// GetBucketCapacity returns the rate limit for the rate limiter.
func (gtb *GlobalTokenBucket) GetBucketCapacity() float64 {
	gtb.mu.RLock()
	defer gtb.mu.RUnlock()
	return gtb.bucketCapacity
}

// SetFillAmount sets the default fill amount for the rate limiter.
func (gtb *GlobalTokenBucket) SetFillAmount(fillAmount float64) {
	gtb.mu.Lock()
	defer gtb.mu.Unlock()
	gtb.fillAmount = fillAmount
}

// Name returns the name of the DistCacheRateTracker.
func (gtb *GlobalTokenBucket) Name() string {
	return gtb.name
}

// Close cleans up DMap held within the DistCacheRateTracker.
func (gtb *GlobalTokenBucket) Close() error {
	gtb.mu.Lock()
	defer gtb.mu.Unlock()
	err := gtb.dc.DeleteDMap(gtb.name)
	if err != nil {
		return err
	}
	return nil
}

// TakeIfAvailable increments value in label by n and returns whether n events should be allowed along with the remaining value (limit - new n) after increment and the current count for the label.
// If an error occurred it returns true, 0 and 0 (fail open).
func (gtb *GlobalTokenBucket) TakeIfAvailable(label string, n float64) (bool, float64, float64) {
	if gtb.GetPassThrough() {
		return true, 0, 0
	}

	if gtb.GetBucketCapacity() == 0 {
		return false, 0, 0
	}

	req := takeNRequest{
		Want:    n,
		CanWait: false,
	}
	// encode request
	reqBytes, err := utils.MarshalGob(req)
	if err != nil {
		log.Autosample().Errorf("error encoding request: %v", err)
		return true, 0, 0
	}

	resultBytes, err := gtb.dMap.Function(label, TakeNFunction, reqBytes)
	if err != nil {
		log.Autosample().Error().Err(err).Str("dmapName", gtb.dMap.Name()).Msg("error taking from token bucket")
		return true, 0, 0
	}

	var resp takeNResponse

	err = utils.UnmarshalGob(resultBytes, &resp)
	if err != nil {
		log.Autosample().Errorf("error decoding response: %v", err)
		return true, 0, 0
	}

	return resp.Ok, resp.Remaining, resp.Current
}

// Take increments value in label by n and returns whether n events should be allowed along with the remaining value (limit - new n) after increment and the current count for the label.
// It also returns the wait time at which the tokens will be available.
func (gtb *GlobalTokenBucket) Take(label string, n float64) (bool, time.Duration, float64, float64) {
	if gtb.GetPassThrough() {
		return true, 0, 0, 0
	}

	if gtb.GetBucketCapacity() == 0 {
		return false, 0, 0, 0
	}

	req := takeNRequest{
		Want:    n,
		CanWait: true,
	}
	// encode request
	reqBytes, err := utils.MarshalGob(req)
	if err != nil {
		log.Autosample().Errorf("error encoding request: %v", err)
		return true, 0, 0, 0
	}

	resultBytes, err := gtb.dMap.Function(label, TakeNFunction, reqBytes)
	if err != nil {
		log.Autosample().Error().Err(err).Str("dmapName", gtb.dMap.Name()).Msg("error taking from token bucket")
		return true, 0, 0, 0
	}

	var resp takeNResponse

	err = utils.UnmarshalGob(resultBytes, &resp)
	if err != nil {
		log.Autosample().Errorf("error decoding response: %v", err)
		return true, 0, 0, 0
	}
	var waitTime time.Duration
	if !resp.AvailableAt.IsZero() {
		waitTime = time.Until(resp.AvailableAt)
		if waitTime < 0 {
			waitTime = 0
		}
	}

	return resp.Ok, waitTime, resp.Remaining, resp.Current
}

// Return returns n tokens to the bucket.
func (gtb *GlobalTokenBucket) Return(label string, n float64) (float64, float64) {
	_, remaining, current := gtb.TakeIfAvailable(label, -n)
	return remaining, current
}

// Per-label tracking in distributed cache.
type tokenBucketState struct {
	LastFill  time.Time
	Available float64
}

type takeNRequest struct {
	Want    float64
	CanWait bool
}

type takeNResponse struct {
	AvailableAt time.Time
	Current     float64
	Remaining   float64
	Ok          bool
}

// takeN takes a number of tokens from the bucket.
func (gtb *GlobalTokenBucket) takeN(key string, stateBytes, argBytes []byte) ([]byte, []byte, error) {
	gtb.mu.RLock()
	defer gtb.mu.RUnlock()

	// Decode currentState from gob encoded currentStateBytes
	now := time.Now()

	state, err := gtb.fastForwardState(now, stateBytes)
	if err != nil {
		return nil, nil, err
	}

	// Decode arg from gob encoded argBytes
	var arg takeNRequest
	if argBytes != nil {
		err = utils.UnmarshalGob(argBytes, &arg)
		if err != nil {
			log.Autosample().Errorf("error decoding arg: %v", err)
			return nil, nil, err
		}
	}

	result := takeNResponse{
		Ok:          true,
		AvailableAt: now,
	}

	state.Available -= arg.Want

	if arg.Want > 0 {
		if state.Available < 0 {
			if gtb.fillAmount != 0 {
				waitTime := time.Duration(-state.Available / gtb.fillAmount * float64(gtb.interval))
				availableAt := now.Add(waitTime)
				result.AvailableAt = availableAt
			}
			result.Ok = arg.CanWait
			// return the tokens to the bucket if the request is not ok
			if !result.Ok {
				state.Available += arg.Want
			}
		}
	}

	if state.Available > gtb.bucketCapacity {
		state.Available = gtb.bucketCapacity
	}

	result.Remaining = state.Available
	result.Current = gtb.bucketCapacity - state.Available

	// Encode result to gob encoded resultBytes
	resultBytes, err := utils.MarshalGob(result)
	if err != nil {
		log.Autosample().Errorf("error encoding result: %v", err)
		return nil, nil, err
	}

	// Encode currentState to gob encoded newStateBytes
	newStateBytes, err := utils.MarshalGob(state)
	if err != nil {
		log.Autosample().Errorf("error encoding new state: %v", err)
		return nil, nil, err
	}

	return newStateBytes, resultBytes, nil
}

func (gtb *GlobalTokenBucket) fastForwardState(now time.Time, stateBytes []byte) (*tokenBucketState, error) {
	var state tokenBucketState

	if stateBytes != nil {
		err := utils.UnmarshalGob(stateBytes, &state)
		if err != nil {
			log.Autosample().Errorf("error decoding current state: %v", err)
			return nil, err
		}
	} else {
		state.LastFill = now
		state.Available = gtb.bucketCapacity
	}

	if gtb.fillAmount != 0 {
		// Calculate the time passed since the last fill
		sinceLastFill := now.Sub(state.LastFill)
		fillAmount := 0.0
		if gtb.continuousFill {
			fillAmount = gtb.fillAmount * float64(sinceLastFill) / float64(gtb.interval)
			state.LastFill = now
		} else if sinceLastFill >= gtb.interval {
			fills := int(sinceLastFill / gtb.interval)
			if fills > 0 {
				fillAmount = gtb.fillAmount * float64(fills)
				state.LastFill = state.LastFill.Add(time.Duration(fills) * gtb.interval)
			}
		}
		// Fill the calculated amount
		state.Available += fillAmount

		if state.Available > gtb.bucketCapacity {
			state.Available = gtb.bucketCapacity
		}
	}
	return &state, nil
}

// SetPassThrough sets the pass through flag.
func (gtb *GlobalTokenBucket) SetPassThrough(passThrough bool) {
	gtb.mu.Lock()
	defer gtb.mu.Unlock()
	gtb.passThrough = passThrough
}

// GetPassThrough returns the pass through flag.
func (gtb *GlobalTokenBucket) GetPassThrough() bool {
	gtb.mu.RLock()
	defer gtb.mu.RUnlock()
	return gtb.passThrough
}

// Make sure TokenBucketRateTracker implements Limiter interface.
var _ ratelimiter.RateLimiter = (*GlobalTokenBucket)(nil)
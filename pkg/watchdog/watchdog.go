// Ported from - https://github.com/raulk/go-watchdog
package watchdog

import (
	"context"
	"errors"
	"fmt"
	"math"
	"runtime"
	"runtime/debug"

	"github.com/elastic/gosigar"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"

	watchdogv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/common/watchdog/v1"
	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/jobs"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/panichandler"
	"github.com/fluxninja/aperture/pkg/status"
)

// swagger:operation POST /watchdog common-configuration Watchdog
// ---
// x-fn-config-env: true
// parameters:
// - name: memory
//   in: body
//   schema:
//     "$ref": "#/definitions/WatchdogConfig"
// - name: scheduler
//   in: body
//   schema:
//     "$ref": "#/definitions/JobGroupConfig"

const (
	watchdogConfigKey    = "watchdog.memory"
	watchdogGroup        = "watchdog"
	watchdogMultiJobName = "service"
	heapStatusKey        = "watchdog.heap"
	// PolicyTempDisabled is a marker value for policies to signal that the policy
	// is temporarily disabled. Use it when all hope is lost to turn around from
	// significant memory pressure (such as when above an "extreme" watermark).
	policyTempDisabled uint64 = math.MaxUint64
)

// Module is a fx module that provides annotated Watchdog jobs and triggers Watchdog checks.
func Module() fx.Option {
	return fx.Options(
		jobs.JobGroupConstructor{Name: watchdogGroup}.Annotate(),
		jobs.MultiJobConstructor{Name: watchdogMultiJobName, JobGroupName: watchdogGroup}.Annotate(),
		fx.Invoke(Constructor{Key: watchdogConfigKey}.setupWatchdog),
	)
}

// Constructor holds fields to set up the Watchdog.
type Constructor struct {
	// Key for config
	Key string
	// Default config
	DefaultConfig WatchdogConfig
}

// WatchdogIn holds parameters for setupWatchdog.
type WatchdogIn struct {
	fx.In

	StatusRegistry *status.Registry
	JobGroup       *jobs.JobGroup `name:"watchdog"`
	WatchdogJob    *jobs.MultiJob `name:"watchdog.service"`
	Unmarshaller   config.Unmarshaller
	Lifecycle      fx.Lifecycle
}

func (constructor Constructor) setupWatchdog(in WatchdogIn) error {
	config := constructor.DefaultConfig

	if err := in.Unmarshaller.UnmarshalKey(constructor.Key, &config); err != nil {
		log.Error().Err(err).Msg("Unable to deserialize watchdog policy!")
		return err
	}

	var gcs *gcSentinel
	// Heap memory check
	var hp *heapPolicy

	in.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			var err error

			// CGroup memory check
			if runtime.GOOS == "linux" {
				job := &jobs.BasicJob{
					JobBase: jobs.JobBase{
						JobName: "cgroup",
					},
				}
				if config.CGroup.WatermarksPolicy.Enabled {
					cgw := &cgroupWatermarks{WatermarksPolicy: config.CGroup.WatermarksPolicy}
					job.JobFunc = cgw.Check
				} else if config.CGroup.AdaptivePolicy.Enabled {
					cga := &cgroupAdaptive{AdaptivePolicy: config.CGroup.AdaptivePolicy}
					job.JobFunc = cga.Check
				}
				err = in.WatchdogJob.RegisterJob(job)
				if err != nil {
					return err
				}
			}

			if config.System.WatermarksPolicy.Enabled || config.System.AdaptivePolicy.Enabled {
				job := &jobs.BasicJob{
					JobBase: jobs.JobBase{
						JobName: "system",
					},
				}
				// System memory check
				if config.System.WatermarksPolicy.Enabled {
					sw := &systemWatermarks{WatermarksPolicy: config.System.WatermarksPolicy}
					job.JobFunc = sw.Check
				} else if config.System.AdaptivePolicy.Enabled {
					sa := &systemAdaptive{AdaptivePolicy: config.System.AdaptivePolicy}
					job.JobFunc = sa.Check
				}
				err = in.WatchdogJob.RegisterJob(job)
				if err != nil {
					return err
				}
			}

			if config.Heap.WatermarksPolicy.Enabled || config.Heap.AdaptivePolicy.Enabled {
				hp = newHeapPolicy(config.Heap)
				s := status.NewStatus(nil, nil)
				err := in.StatusRegistry.Push(heapStatusKey, s)
				if err != nil {
					return err
				}
			}

			gcs = newSentinel()
			// start a go routine to track GC
			panichandler.Go(func() {
				for {
					select {
					case <-gcs.gcTriggered:
						log.Trace().Msg("GC detected, triggering watchdog checks")
						in.JobGroup.TriggerJob(watchdogMultiJobName)
						if hp != nil {
							details, e := hp.checkHeap()
							s := status.NewStatus(details, e)
							err := in.StatusRegistry.Push(heapStatusKey, s)
							if err != nil {
								log.Error().Err(err).Msg("Unable to push heap check results to status registry")
							}
						}
					case <-gcs.ctx.Done():
						return
					}
				}
			})

			return nil
		},
		OnStop: func(context.Context) error {
			gcs.stop()
			_ = in.WatchdogJob.DeregisterJob("cgroup")
			_ = in.WatchdogJob.DeregisterJob("system")
			in.StatusRegistry.Delete(heapStatusKey)
			return nil
		},
	})

	return nil
}

// GC Sentinel trigger.
type gcSentinel struct {
	gcTriggered chan struct{}
	ctx         context.Context
	cancel      context.CancelFunc
}

func newSentinel() *gcSentinel {
	gcs := &gcSentinel{}
	gcs.gcTriggered = make(chan struct{}, 16)
	gcs.ctx, gcs.cancel = context.WithCancel(context.Background())

	// this non-zero sized struct is used as a sentinel to detect when a GC
	// run has finished, by setting and resetting a finalizer on it.
	// it essentially creates a GC notification "flywheel"; every GC will
	// trigger this finalizer, which will reset itself so it gets notified
	// of the next GC, breaking the cycle when the Watchdog is stopped.
	type sentinel struct{ a *int }
	var finalizer func(o *sentinel)
	finalizer = func(o *sentinel) {
		// reset so it triggers on the next GC.
		runtime.SetFinalizer(o, finalizer)

		select {
		case <-gcs.ctx.Done():
			return
		default:
		}

		select {
		case gcs.gcTriggered <- struct{}{}:
		default:
			log.Warn().Msg("Failed to queue GC trigger, channel backlogged")
		}
	}

	runtime.SetFinalizer(&sentinel{a: nil}, finalizer) // start the flywheel.
	return gcs
}

func (gcs *gcSentinel) stop() {
	gcs.cancel()
}

/* System policies */

func systemUsage() (uint64, uint64, error) {
	var sysmem gosigar.Mem
	if err := (*gosigar.Mem).Get(&sysmem); err != nil {
		return 0, 0, fmt.Errorf("failed to get system memory stats: %w", err)
	}
	return sysmem.Total, sysmem.ActualUsed, nil
}

type systemWatermarks struct {
	WatermarksPolicy
}

// Check evaluates the system memory usage and runs GC at configured watermarks of memory utilization.
func (policy *systemWatermarks) Check(ctx context.Context) (proto.Message, error) {
	log.Debug().Msg("System watermarks check triggered")
	return check(policy, ctx, systemUsage)
}

type systemAdaptive struct {
	AdaptivePolicy
}

// Check evaluates the system memory usage and runs GC at configured adaptive thresholds of memory utilization.
func (policy *systemAdaptive) Check(ctx context.Context) (proto.Message, error) {
	log.Debug().Msg("System adaptive check triggered")
	return check(policy, ctx, systemUsage)
}

// Heap Policy.
type heapPolicy struct {
	HeapConfig
	originalGoGC int
	currGoGC     int
}

func newHeapPolicy(config HeapConfig) *heapPolicy {
	hp := heapPolicy{HeapConfig: config}

	// get the initial effective GoGC; guess it's 100 (default), and restore
	// it to whatever it actually was. This works because SetGCPercent
	// returns the previous value.
	hp.originalGoGC = debug.SetGCPercent(debug.SetGCPercent(100))
	hp.currGoGC = hp.originalGoGC
	return &hp
}

func (hp *heapPolicy) checkHeap() (proto.Message, error) {
	log.Debug().Msg("Heap check triggered")
	if hp.Limit == 0 {
		return nil, fmt.Errorf("cannot use zero limit for heap-driven watchdog")
	}

	var err error
	var threshold uint64
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	// heapMarked is the amount of heap that was marked as live by GC.
	// it is inferred from our current GoGC and the new target picked.
	heapMarked := uint64(float64(memstats.NextGC) / (1 + float64(hp.currGoGC)/100))
	usage := memstats.HeapAlloc
	switch {
	case hp.HeapConfig.WatchdogPolicyType.WatermarksPolicy.Enabled:
		threshold = hp.HeapConfig.WatchdogPolicyType.WatermarksPolicy.nextThreshold(hp.Limit, usage)
	case hp.HeapConfig.WatchdogPolicyType.AdaptivePolicy.Enabled:
		threshold = hp.HeapConfig.WatchdogPolicyType.AdaptivePolicy.nextThreshold(hp.Limit, usage)
	default:
		log.Panic().Msg("checkHeap called on disabled policy")
	}

	// calculate how much to set GoGC to honor the next trigger point.
	// next=PolicyTempDisabled value would make currGoGC extremely high,
	// greater than originalGoGC, and therefore we'd restore originalGoGC.
	hp.currGoGC = int(((float64(threshold) / float64(heapMarked)) - float64(1)) * 100)
	if hp.currGoGC >= hp.originalGoGC {
		hp.currGoGC = hp.originalGoGC
	} else if hp.currGoGC < hp.MinGoGC {
		err = errors.New("heap driven watchdog reached minimum threshold for GoGC value")
		// cap GoGC to avoid overscheduling.
		hp.currGoGC = hp.MinGoGC
	}

	debug.SetGCPercent(hp.currGoGC)
	runtime.ReadMemStats(&memstats)

	if threshold == policyTempDisabled {
		threshold = memstats.NextGC
	}
	log.Info().
		Uint64("heap_alloc", memstats.HeapAlloc).
		Uint64("heap_marked", heapMarked).
		Uint64("memstats_nextGC", memstats.NextGC).
		Uint64("threshold", threshold).
		Int("current_GOGC", hp.currGoGC).
		Msg("GC finished")

	result := &watchdogv1.HeapResult{
		Limit:        hp.Limit,
		HeapMarked:   heapMarked,
		Threshold:    threshold,
		CurrGogc:     int32(hp.currGoGC),
		OriginalGogc: int32(hp.originalGoGC),
		TotalAlloc:   memstats.TotalAlloc,
		Sys:          memstats.Sys,
		Mallocs:      memstats.Mallocs,
		Frees:        memstats.Frees,
		HeapAlloc:    memstats.HeapAlloc,
		HeapSys:      memstats.HeapSys,
		HeapInuse:    memstats.HeapInuse,
		HeapReleased: memstats.HeapReleased,
		HeapObjects:  memstats.HeapObjects,
		NextGc:       memstats.NextGC,
		LastGc:       memstats.LastGC,
		PauseTotalNs: memstats.PauseTotalNs,
		NumGc:        memstats.NumGC,
		NumForcedGc:  memstats.NumForcedGC,
	}

	return result, err
}

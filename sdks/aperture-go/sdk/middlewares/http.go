package middlewares

import (
	"fmt"
	"net/http"
	"regexp"

	aperture "github.com/fluxninja/aperture-go/v2/sdk"
)

// HTTPMiddleware is the interface for the HTTP middleware
type HTTPMiddleware interface {
	Handle(http.Handler) http.Handler
}

type httpMiddleware struct {
	client       aperture.Client
	controlPoint string
	ignoredPaths *[]regexp.Regexp
}

// NewHTTPMiddleware creates a new HTTPMiddleware struct
func NewHTTPMiddleware(client aperture.Client, controlPoint string, ignoredPaths *[]regexp.Regexp) HTTPMiddleware {
	return &httpMiddleware{
		client:       client,
		controlPoint: controlPoint,
		ignoredPaths: ignoredPaths,
	}
}

// Handle takes a http.Handler and returns a new http.Handler with the middleware applied
func (m *httpMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the path is ignored, skip the middleware
		if m.ignoredPaths != nil {
			for _, ignoredPath := range *m.ignoredPaths {
				if ignoredPath.MatchString(r.URL.Path) {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		req := PrepareCheckHTTPRequestForHTTP(r, m.client.GetLogger(), m.controlPoint)

		flow, err := m.client.StartHTTPFlow(r.Context(), req)
		if err != nil {
			m.client.GetLogger().Info("Aperture flow control got error. Returned flow defaults to Allowed.", "flow.Accepted()", flow.Accepted())
		}

		defer func() {
			// Need to call End() on the Flow in order to provide telemetry to Aperture Agent for completing the control loop.
			// The first argument captures whether the feature captured by the Flow was successful or resulted in an error.
			// The second argument is error message for further diagnosis.
			err := flow.End(aperture.OK)
			if err != nil {
				m.client.GetLogger().Info("Aperture flow control end got error.", "error", err)
			}
		}()

		if flow.Accepted() {
			// Simulate work being done
			next.ServeHTTP(w, r)
		} else {
			resp := flow.CheckResponse().GetDeniedResponse()
			// If there was connection error, the response will be nil.
			if resp == nil {
				w.WriteHeader(http.StatusServiceUnavailable)
			} else {
				w.WriteHeader(int(resp.GetStatus()))
				for key, value := range resp.GetHeaders() {
					w.Header().Set(key, value)
				}
				_, err := fmt.Fprint(w, resp.GetBody())
				if err != nil {
					m.client.GetLogger().Info("Aperture flow control respond body got an error.", "error", err)
				}
			}
		}
	})
}
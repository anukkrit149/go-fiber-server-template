package heimdall

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	env                 string
	requestCounter      *prometheus.CounterVec
	requestDurations    *prometheus.HistogramVec
	requestDurationsSum *prometheus.SummaryVec
)

func init() {
	env = os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// metric for client http request count.
	requestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "name_client_http_requests_count",
			Help: "Total HTTP request via clients",
		},
		[]string{"env", "code", "host", "method", "url", "reason"},
	)

	// metric for client http request latency.
	requestDurations = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "name_client_http_request_duration_ms_hist",
			Help:    "A histogram of client request latencies.",
			Buckets: []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096},
		},
		[]string{"env", "code", "host", "method", "uri"},
	)

	// summary for client http request latency.
	requestDurationsSum = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "name_client_http_req_durations_ms_summary",
			Help:       "Client HTTP latency distributions summary.",
			Objectives: map[float64]float64{0.5: 0.5, 0.9: 0.9, 0.99: 0.99},
		},
		[]string{"env", "code", "method", "host", "uri"},
	)
}

// RoundTripperFunc
type RoundTripperFunc func(req *http.Request) (*http.Response, error)

// RoundTrip implements the RoundTripper interface.
func (rt RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return rt(r)
}

// InstrumentRoundTripper instruments http calls
func InstrumentRoundTripper(next http.RoundTripper) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		start := time.Now()
		resp, err := next.RoundTrip(r)
		durationMs := float64(time.Since(start).Milliseconds())
		statusCode := "unknown"
		reason := ""

		if resp != nil {
			statusCode = fmt.Sprintf("%v", resp.StatusCode)
		}

		if err == nil {
			requestDurations.WithLabelValues(
				env,
				statusCode,
				r.Host,
				r.Method,
				r.URL.String(),
			).Observe(durationMs)

			requestDurationsSum.WithLabelValues(
				env,
				statusCode,
				r.Method,
				r.Host,
				r.URL.String(),
			).Observe(durationMs)
		} else {
			reason = err.Error()
		}

		requestCounter.WithLabelValues(
			env,
			statusCode,
			r.Host,
			r.Method,
			r.URL.String(),
			reason,
		).Inc()

		return resp, err
	}
}

package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for handler in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration, requestCounter)
}

// MetricsMiddleware instruments requests with Prometheus metrics
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// wrap response writer to capture status
		rw := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration)
		requestCounter.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(rw.statusCode)).Inc()
	})
}

// statusRecorder captures HTTP status codes
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

// PrometheusHandler returns the Prometheus handler for /metrics
func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}

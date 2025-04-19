package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// LoggingMiddleware logs method, path and duration of each request.
func LoggingMiddleware(logger *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Infow("HTTP request",
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(start),
			)
		})
	}
}

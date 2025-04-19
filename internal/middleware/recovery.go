package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"
)

// RecoveryMiddleware recovers from panics and logs the stack trace.
func RecoveryMiddleware(logger *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Errorw("panic recovered",
						"error", err,
						"stack", string(debug.Stack()),
					)
					http.Error(w, fmt.Sprintf("internal server error: %v", err), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

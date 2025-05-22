package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func LoggingMiddleware(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)

			log.Info("HTTP Request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.Path),
				zap.String("remote", r.RemoteAddr),
				zap.Duration("duration", duration),
			)
		})
	}
}

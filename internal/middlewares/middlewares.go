package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestStats(next http.Handler, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := time.Now()
		next.ServeHTTP(w, r)
		log.Info("request", "method", r.Method, "url", r.URL.String(), "time(nanoseconds)", time.Since(current).Nanoseconds())
	})
}

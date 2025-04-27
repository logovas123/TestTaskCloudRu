package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// мидлвар выводит информацию о пришедшем запросе
func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		slog.Info("get new request",
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
			"url", r.URL.String(),
			"time", time.Since(start),
		)
	})
}

package serverpool

import (
	"log/slog"
	"net/http"
)

// главный обработчик запроса к балансировщику нагрузки
func (s *ServerPool) LBHandler(w http.ResponseWriter, r *http.Request) {
	conn := s.GetNextActiveConn()
	if conn != nil {
		slog.Info("find live connect", "url", conn.GetURL().String())
		conn.GetReverseProxy().ServeHTTP(w, r)
		slog.Info("request send success", "url", conn.GetURL().String())
		return
	}

	slog.Error("all servers dead")
	http.Error(w, "service not available", http.StatusServiceUnavailable)
}

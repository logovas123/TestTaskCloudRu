package serverpool

import (
	"log/slog"
	"net"
	"net/url"
	"time"
)

// проверяет все сервера на жизнь
func (s *ServerPool) HealthCheck() {
	for _, b := range s.listOfSrvs {
		status := "alive"
		alive := isSrvNodeAlive(b.GetURL())
		b.SetAlive(alive)
		if !alive {
			status = "dead"
		}
		slog.Info("server health check", "type", "HEALTH_CHECK", "url", b.GetURL().String(), "status", status)
	}
}

// проверяет сервер на жизнь - отправляет запрос с таймаутом
func isSrvNodeAlive(u *url.URL) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

package lbserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lb-service/config"
	serverpool "lb-service/internal/serverPool"
	srvnode "lb-service/internal/srvNode"
)

type ServerPoolInterface interface {
	LBHandler(w http.ResponseWriter, r *http.Request)
	AddSrvToList(srv serverpool.ServerNodeHandler)
	HealthCheck()
}

type Server struct {
	cfg  *config.Config
	srv  *http.Server
	pool ServerPoolInterface
}

func NewLoadBalancerServer(cfg *config.Config, pool ServerPoolInterface, mux http.Handler) *Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: mux,
	}

	return &Server{
		cfg:  cfg,
		srv:  srv,
		pool: pool,
	}
}

// запуск сервера
func (s *Server) Run() error {
	for _, urlString := range s.cfg.ListOfSrvs {
		urlServer, err := url.Parse(urlString)
		if err != nil {
			slog.Error("url no valid", "error", err)
			return err
		}

		proxy := httputil.NewSingleHostReverseProxy(urlServer)

		srv := srvnode.NewSrvNode(urlServer, proxy)

		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			slog.Error("proxy error", "error", err, "url", urlServer.String())

			srv.SetAlive(false)

			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		}

		s.pool.AddSrvToList(srv)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		t := time.NewTicker(30 * time.Second) // каждые 30 сек проверка на жизнь
		defer t.Stop()

		for {
			select {
			case <-t.C:
				slog.Info("health check...")
				s.pool.HealthCheck()
				slog.Info("health check completed")
			case <-ctx.Done():
				slog.Info("stopping health check")
				return
			}
		}
	}()

	go func() {
		slog.Info("server start", "port", s.cfg.Port)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error starting server",
				"error", err,
				"port", s.cfg.Port,
			)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutdown server...")

	err := s.srv.Shutdown(ctx)
	if err != nil {
		slog.Error("can't server shutdown", "error", err)
	} else {
		slog.Info("server closed success")
	}

	return nil
}

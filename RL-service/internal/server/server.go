package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rl-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	cfg  *config.Config
	pool *pgxpool.Pool
	srv  *http.Server
}

func NewServer(cfg *config.Config, pool *pgxpool.Pool) *Server {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", cfg.URLServer.Port),
	}

	return &Server{
		cfg:  cfg,
		pool: pool,
		srv:  srv,
	}
}

// метод запускает сервер
func (s *Server) Run() error {
	slog.Info("Running server...")

	s.MapHandlers()
	slog.Info("router create success")

	go func() {
		slog.Info("server start", "port", s.cfg.URLServer.Port)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error starting server",
				"error", err,
				"port", s.cfg.URLServer.Port,
			)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.srv.Shutdown(ctx)
	if err != nil {
		slog.Error("can't server shutdown", "error", err)
	} else {
		slog.Info("server closed success")
	}

	slog.Info("closing postgres pool conns...")
	s.pool.Close()
	slog.Info("pool conns closed success")

	return nil
}

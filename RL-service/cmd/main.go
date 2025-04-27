package main

import (
	"context"
	"log/slog"
	"os"

	"rl-service/internal/config"
	"rl-service/internal/db"
	"rl-service/internal/server"
	"rl-service/pkg/logger"
)

func main() {
	logger.Init()
	slog.Info("logger init success")

	cfg := config.MustLoad()
	slog.Info("config success load")

	ctx := context.Background()

	pool, err := db.NewConnPostgres(ctx, cfg.DB)
	if err != nil {
		slog.Error("error create coonect to db:",
			"error", err,
		)
		os.Exit(1)
	}

	s := server.NewServer(cfg, pool)
	slog.Info("server create success")
	if err := s.Run(); err != nil {
		slog.Error("can't start server", "error", err)
		os.Exit(1)
	}

	slog.Info("server gracefull shutdown")
}

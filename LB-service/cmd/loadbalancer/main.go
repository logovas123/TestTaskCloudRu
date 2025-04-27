package main

import (
	"flag"
	"log/slog"
	"os"

	"lb-service/config"
	"lb-service/internal/lbserver"
	serverpool "lb-service/internal/serverPool"
	"lb-service/pkg/logger"
)

func main() {
	logger.Init()
	slog.Info("logger init success")

	configPath := flag.String("config", "configs/config.json", "Path to config file")
	flag.Parse()
	slog.Info("flag parse success")

	cfg, err := config.MustLoad(*configPath) // загружаем конфиг (json)
	if err != nil {
		slog.Error("can't load config", "error", err)
		os.Exit(1)
	}
	slog.Info("config load success")

	pool := serverpool.NewServerPool() // создаём пул серверов
	slog.Info("pool create success")

	mux := lbserver.NewRouter(pool) // создаём роутер
	slog.Info("router create success")

	srv := lbserver.NewLoadBalancerServer(cfg, pool, mux)
	if err := srv.Run(); err != nil {
		slog.Error("can't start server", "error", err)
		os.Exit(1)
	}
}

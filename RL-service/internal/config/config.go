package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB        *DBConfig   `envconfig:"DB" required:"true"`
	URLServer *PathConfig `envconfig:"URL" required:"true"`
}

// загружаем конфиг из env файла
func MustLoad() *Config {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("error load env file",
			"error", err,
		)
		os.Exit(1)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		slog.Error("error from envconfig",
			"error", err,
		)
		os.Exit(1)
	}

	slog.Info("env file loaded")

	return &cfg
}

package config

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
)

type Config struct {
	Port       string   `json:"port"`
	ListOfSrvs []string `json:"list_of_srvs"`
}

// метод парсит конфиг
func MustLoad(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("error opening config.json", "error", err)
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		slog.Error("error decoding config.json", "error", err)
		return nil, err
	}

	if len(cfg.ListOfSrvs) == 0 {
		slog.Error("list of servers empty")
		return nil, errors.New("list of servers empty")
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return &cfg, nil
}

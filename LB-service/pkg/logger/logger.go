package logger

import (
	"log/slog"
	"os"
)

// инициализируем логгер
func Init() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	slog.SetDefault(log)
}

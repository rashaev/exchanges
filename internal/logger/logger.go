package logger

import (
	"exchanges/internal/config"
	"log/slog"
	"os"
)

type Level int

const (
	Debug Level = -4
	Info  Level = 0
	Warn  Level = 4
	Error Level = 8
)

func New(cfg *config.Config) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.Logger.Level,
	})

	return slog.New(handler)
}

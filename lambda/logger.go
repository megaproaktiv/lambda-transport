package lambda

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

const (
	LevelDebug = slog.Level(-4)
	LevelInfo  = slog.Level(0)
	LevelWarn  = slog.Level(4)
	LevelError = slog.Level(8)
)

func init() {
	// LOGGER
	handlerInfo := slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{Level: LevelInfo})
	Logger = slog.New(handlerInfo)
}

func SetLogLevelDebug() {
	handlerDebug := slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{Level: LevelDebug})
	Logger = slog.New(handlerDebug)
}

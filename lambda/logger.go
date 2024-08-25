package lambda

import (
	"log/slog"
)

var Logger *slog.Logger

const (
	LevelDebug = slog.Level(-4)
	LevelInfo  = slog.Level(0)
	LevelWarn  = slog.Level(4)
	LevelError = slog.Level(8)
)

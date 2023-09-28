package logger

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

var (
	Info  = logger.Info
	Warn  = logger.Warn
	Error = logger.Error
)

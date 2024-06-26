package utils

import (
	"golang.org/x/exp/slog"
	"os"
)

func InitLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func GetLogger() *slog.Logger {
	return slog.Default()
}

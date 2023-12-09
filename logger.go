package main

import (
	"log/slog"
	"os"

	"github.com/drornir/cloudex/pkg/config"
)

func newLogger(conf config.Config) *slog.Logger {
	var logLevel slog.Level
	if conf.LogLevel != "" {
		if err := logLevel.UnmarshalText([]byte(conf.LogLevel)); err != nil {
			logLevel = slog.LevelInfo
		}
	}

	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		// AddSource: true,
		Level: logLevel,
	}))
}

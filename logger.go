package main

import (
	"log/slog"
	"os"
	"sync/atomic"

	"github.com/drornir/cloudex/pkg/config"
)

func newLogger(conf config.Config) *slog.Logger {
	var logLevel slog.Level
	if conf.LogLevel != "" {
		if err := logLevel.UnmarshalText([]byte(conf.LogLevel)); err != nil {
			logLevel = slog.LevelInfo
		}
	}

	var logCounter uint64

	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		// AddSource: true,
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "time":
				c := atomic.AddUint64(&logCounter, 1)
				return slog.Uint64("time", c)
			default:
				return a
			}
		},
	}))
}

package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once
var logger *zap.Logger

// InitLogger returns logger for a given log level
func InitLogger() *zap.Logger {
	once.Do(func() {
		atom := zap.NewAtomicLevel()
		l, err := zapcore.ParseLevel("debug")
		if err != nil {
			fmt.Printf("err: %#v", err)
		}

		cfg := zap.Config{
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: true,
			Encoding:          "json",
			EncoderConfig: zapcore.EncoderConfig{
				CallerKey:     "caller",
				LevelKey:      "level",
				MessageKey:    "msg",
				StacktraceKey: "stacktrace",
				TimeKey:       "ts",
				EncodeCaller:  zapcore.ShortCallerEncoder,
				EncodeLevel:   zapcore.LowercaseLevelEncoder,
				EncodeTime:    zapcore.ISO8601TimeEncoder,
			},
			ErrorOutputPaths: []string{"stderr"},
			Level:            atom,
			OutputPaths:      []string{"stdout"},
		}

		logger = zap.Must(cfg.Build())
		atom.SetLevel(l)

	})
	return logger
}

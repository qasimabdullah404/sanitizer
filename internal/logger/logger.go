package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level string) *zap.Logger {
	lvl := parseLevel(level)
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(lvl)
	config.Encoding = "json"
	config.OutputPaths = []string{"stdout"}

	logger, _ := config.Build()
	return logger
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

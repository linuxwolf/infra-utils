package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogging(verbosity int) *zap.SugaredLogger {
	var level zapcore.Level
	switch verbosity {
	case 0:
		level = zap.WarnLevel
	case 1:
		level = zap.InfoLevel
	default:
		level = zap.DebugLevel
	}

	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(level)
	cfg.Encoding = "console"

	logger, _ := cfg.Build()
	return logger.Sugar()
}

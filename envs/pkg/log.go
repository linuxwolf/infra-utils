package pkg

import "go.uber.org/zap"

func SetupLogging() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.WarnLevel)
	cfg.Encoding = "console"

	logger, _ := cfg.Build()
	return logger.Sugar()
}

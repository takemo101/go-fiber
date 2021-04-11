package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger structure
type Logger struct {
	Zap *zap.SugaredLogger
}

// NewLogger sets up logger
func NewLogger(config Config) Logger {

	zapConfig := zap.NewDevelopmentConfig()

	if config.App.Debug {
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if config.Log.Server != "" {
		zapConfig.OutputPaths = []string{config.Log.Server}
	}

	logger, _ := zapConfig.Build()

	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}
}

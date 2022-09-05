package log

import (
	defaultLog "log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup() {
	log, err := newLogger()
	if err != nil {
		defaultLog.Fatalf("failure to initialize logger: %v", err)
	}

	_ = zap.ReplaceGlobals(log)
}

func newLogger() (*zap.Logger, error) {
	env := os.Getenv("ENVIRONMENT")

	var config zap.Config
	if env == "development" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if lvl, exists := os.LookupEnv("LOG_LEVEL"); exists {
		lvl = strings.ToLower(lvl)
		switch lvl {
		case "debug":
			config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		case "info":
			config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case "warn":
			config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		case "error":
			config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		case "panic":
			config.Level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
		case "fatal":
			config.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
		}
	}

	return config.Build()
}

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger returns a new logger with the given options
func NewLogger() *zap.Logger {

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger

}

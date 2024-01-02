package main

import (
	"hello/logger"

	"go.uber.org/zap/zapcore"
)

func main() {
	customLogger := logger.GetLogger()
	customLogger.Info("hello world", zapcore.Field{Key: "name", Type: zapcore.StringType, String: "Jane"})
}

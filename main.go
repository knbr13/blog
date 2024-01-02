package main

import (
	"hello/logger"

	"go.uber.org/zap"
)

func main() {
	customLogger := logger.GetLogger()
	customLogger.Info("hello world", zap.String("Name", "Jane"), zap.Int("Age", 19))
	ll := logger.GetInternalLogger()
	ll.Warn("warning...")
}

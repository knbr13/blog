package main

import (
	"go.uber.org/zap"

	"hello/logger"
)

func main() {
	customLogger := logger.GetLogger()
	customLogger.Info("hello world", zap.String("Name", "Jane"), zap.Int("Age", 19))
	ll := logger.GetInternalLogger()
	ll.Debug("debugging...")
}

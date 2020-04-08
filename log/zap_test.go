package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestInitLogger(t *testing.T) {
	log := InitLogger("test")
	log.Info("Info", zap.String("test", "Info"))
	log.Warn("Warn", zap.String("test", "Warn"))
	log.Debug("Debug", zap.String("test", "Debug"))
	log.Error("Error", zap.String("test", "Error"))
}
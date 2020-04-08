package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestInitLogger(t *testing.T) {
	log := InitLogger("test")
	log.Info("test", zap.String("test", "heihei"))
}
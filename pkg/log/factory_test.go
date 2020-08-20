package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestNewFactory(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("[Error] zap: %v", err)
	}

	fa := NewFactory(logger)
	fa.Bg().Info("test_info", zap.String("level", "info"))
	fa.Bg().Error("test_error", zap.String("level", "error"))
}

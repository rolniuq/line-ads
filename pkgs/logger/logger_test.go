package logger

import (
	"fmt"
	"testing"
)

func TestLoggerInfo(t *testing.T) {
	logger := LoggerMod.Resolve()
	logger.Info("this is info log")
}

func TestLoggerWarning(t *testing.T) {
	params := 100

	logger := LoggerMod.Resolve()
	logger.Warning("this is warning log", params)
}

func TestLoggerError(t *testing.T) {
	logger := LoggerMod.Resolve()
	logger.Error("this is error log", fmt.Errorf("this is error"))
}

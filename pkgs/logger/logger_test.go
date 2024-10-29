package logger

import (
	"fmt"
	"testing"
)

func TestLogger_Error(t *testing.T) {
	logger := LoggerMod.Resolve()

	logger.Error("this is error", fmt.Errorf("panic"))
}

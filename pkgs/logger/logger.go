package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/submodule-org/submodule.go/v2"
)

var LoggerMod = submodule.Make[*Logger](func() *Logger {
	return newLogger()
})

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

type Logger struct {
	logger *log.Logger
}

func newLogger() *Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Println(Red, fmt.Sprintf("[ERROR] %s %v", msg, err))
}

func (l *Logger) Info(msg string) {
	l.logger.Println(Green, msg)
}

func (l *Logger) Warning(msg string, args ...any) {
	l.logger.Println(Yellow, msg, args)
}

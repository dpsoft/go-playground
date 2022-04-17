package logger

import (
	"log"
	"os"
)

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

type BuiltinLogger struct {
	logger *log.Logger
}

func newBuiltinLogger(logger *log.Logger) *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *BuiltinLogger) Debug(args ...any) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Debugf(format string, args ...any) {
	l.logger.Printf(format, args...)
}

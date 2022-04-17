package logger

import (
	"log"
	"os"
)

type Logger interface {
	//Errorf(format string, args ...any)
	//Fatalf(format string, args ...any)
	//Fatal(args ...any)
	//Infof(format string, args ...any)
	//Info(args ...any)
	//Warnf(format string, args ...any)
	Debugf(format string, args ...any)
	Debug(args ...any)
}

type BuiltinLogger struct {
	logger *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *BuiltinLogger) Debug(args ...any) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Debugf(format string, args ...any) {
	l.logger.Printf(format, args...)
}

package logger

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
}

package logz

import (
	"io"
)

type LoggerOption func(Logger) Logger

func WithWriter(w io.Writer) LoggerOption {
	return func(l Logger) Logger {
		return l.Output(w)
	}
}

func WithLevel(lev Level) LoggerOption {
	return func(l Logger) Logger {
		return l.Level(lev)
	}
}

func WithKey(key string) LoggerOption {
	return func(l Logger) Logger {
		Define(key, l)
		return l
	}
}

func SetAsDefault() LoggerOption {
	return func(l Logger) Logger {
		defaultLogger = l
		return defaultLogger
	}
}

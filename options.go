package logz

type LoggerOption func(Logger) Logger

func WithLevel(lev Level) LoggerOption {
	return func(l Logger) Logger {
		return l.Level(lev)
	}
}

func WithKey(key string) LoggerOption {
	return func(l Logger) Logger {
		Append(key, l)
		return l
	}
}

func AppendAs(key string) LoggerOption {
	return func(l Logger) Logger {
		Append(key, l)
		return l
	}
}

func AppendAsDefault() LoggerOption {
	return func(l Logger) Logger {
		rootLogger = l
		return rootLogger
	}
}

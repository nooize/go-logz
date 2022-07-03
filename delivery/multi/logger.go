package multi

import (
	"github.com/nooize/go-logz"
)

// Logger represent structured logger abstraction layer
type multiLogger struct {
	active  bool
	loggers []logz.Logger
}

func (l *multiLogger) Level(level logz.Level) logz.Logger {
	//TODO implement me
	panic("implement me")
}

func (l *multiLogger) GetLevel() logz.Level {
	//TODO implement me
	panic("implement me")
}

func (l *multiLogger) WithLevel(level logz.Level) logz.Event {
	switch level {
	case logz.Trace:
		return l.Trace()
	case logz.Debug:
		return l.Debug()
	case logz.Info:
		return l.Info()
	case logz.Warning:
		return l.Warn()
	case logz.Error:
		return l.Error()
	case logz.Fatal:
		return l.event(logz.Fatal)
	case logz.Disabled:
		return nil
	default:
		return l.event(level)
	}
}

func (l *multiLogger) Trace() logz.Event {
	return l.event(logz.Trace)
}

func (l *multiLogger) Debug() logz.Event {
	return l.event(logz.Debug)
}

func (l *multiLogger) Info() logz.Event {
	return l.event(logz.Info)
}

func (l *multiLogger) Warn() logz.Event {
	return l.event(logz.Warning)
}

func (l *multiLogger) Error() logz.Event {
	return l.event(logz.Error)
}

func (l *multiLogger) Fatal() logz.Event {
	return l.event(logz.Fatal)
}

func (l *multiLogger) event(level logz.Level) logz.Event {
	list := make([]logz.Event, len(l.loggers))
	for i, log := range l.loggers {
		list[i] = log.WithLevel(level)
	}
	l.loggers = list
	return &multiEvent{events: &l.loggers}
}

// Err starts a new message with error level.
func (l *multiLogger) Err(v error) {
	if l.loggers == nil {
		return
	}
	for _, log := range *l.loggers {
		go log.Err(v)
	}
}

package multi

import (
	"github.com/nooize/go-logz"
)

type multiEvent struct {
	active bool
	events []logz.Event
}

func (e *multiEvent) Enabled() bool {
	return m.active
}

func (e *multiEvent) Discard() logz.Event {
	m.active = false
	return m
}

func (e *multiEvent) With(key string, v interface{}) logz.Event {
	if m.loggers == nil {
		return m
	}
	list := make([]logz.Logger, len(*m.loggers))
	for i, log := range *m.loggers {
		list[i] = log.With(key, v)
	}
	*m.loggers = list
}

func (l multiLogger) Send(format string, v ...interface{}) {
	if l.loggers == nil {
		return
	}
	for _, log := range *l.loggers {
		switch lev {
		case logz.Debug:
			go log.Debug(format, v...)
		case logz.Info:
			go log.Info(format, v...)
		case logz.Warning:
			go log.Warn(format, v...)
		case logz.Error:
			go log.Error(format, v...)
		case logz.Fatal:
			go log.Fatal(format, v...)
		case logz.Nop:
		case logz.Disabled:
		}
	}
}

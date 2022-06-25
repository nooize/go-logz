package multi

import (
	"github.com/nooize/go-logz"
)

// Logger represent structured logger abstraction layer
type multiLogger struct {
	active  bool
	loggers []logz.Logger
}

func (m multiLogger) Enabled() bool {
	return m.active
}

func (m multiLogger) Discard() logz.Logger {
	m.active = false
	return m
}

func (m multiLogger) Trace(format string, v ...interface{}) {
	m.send(logz.Trace, format, v...)
}

func (m multiLogger) Debug(format string, v ...interface{}) {
	m.send(logz.Debug, format, v...)
}

func (m multiLogger) Info(format string, v ...interface{}) {
	m.send(logz.Info, format, v...)
}

func (m multiLogger) Warn(format string, v ...interface{}) {
	m.send(logz.Warning, format, v...)
}

func (m multiLogger) Error(format string, v ...interface{}) {
	m.send(logz.Error, format, v...)
}

func (m multiLogger) Fatal(format string, v ...interface{}) {
	m.send(logz.Fatal, format, v...)
}

func (m multiLogger) send(lev logz.Level, format string, v ...interface{}) {
	for _, log := range m.loggers {
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

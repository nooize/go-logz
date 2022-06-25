package sys

import (
	"github.com/nooize/go-logz"
	"log"
)

// Logger represent structured logger abstraction layer
type sysLogger struct {
	level logz.Level
}

func (m sysLogger) Enabled() bool {
	return m.level != logz.Disabled
}

func (m sysLogger) Discard() logz.Logger {
	m.level = logz.Disabled
	return m
}

func (m sysLogger) Trace(format string, v ...interface{}) {
	m.send(logz.Trace, format, v...)
}

func (m sysLogger) Debug(format string, v ...interface{}) {
	m.send(logz.Debug, format, v...)
}

func (m sysLogger) Info(format string, v ...interface{}) {
	m.send(logz.Info, format, v...)
}

func (m sysLogger) Warn(format string, v ...interface{}) {
	m.send(logz.Warning, format, v...)
}

func (m sysLogger) Error(format string, v ...interface{}) {
	m.send(logz.Error, format, v...)
}

func (m sysLogger) Fatal(format string, v ...interface{}) {
	m.send(logz.Fatal, format, v...)
}

func (m sysLogger) send(lev logz.Level, format string, v ...interface{}) {
	if m.level > lev {
		return
	}
	switch lev {
	case logz.Nop:
	case logz.Disabled:
	default:
		log.Printf(format, v...)
	}
}

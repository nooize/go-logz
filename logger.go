package logz

import (
	"log"
)

// Logger represent structured logger abstraction layer
type Logger interface {

	// Level creates a child logger with the minimum accepted level set to level.
	Level(Level) Logger

	// GetLevel returns the current Level of l.
	GetLevel() Level

	// Enabled return false if the Event is going to be filtered out by
	// log level or sampling.
	Enabled() bool

	// Discard disables the event so Send(f) won't print it.
	Discard() Logger

	// Trace starts a new message with trace level.
	//
	// You must call Msg on the returned event in order to send the event.
	Trace(format string, v ...interface{})
	// Debug starts a new message with debug level.
	//
	// You must call Msg on the returned event in order to send the event.
	Debug(format string, v ...interface{})

	// Info starts a new message with info level.
	//
	// You must call Msg on the returned event in order to send the event.
	Info(format string, v ...interface{})

	// Warn starts a new message with warn level.
	//
	// You must call Msg on the returned event in order to send the event.
	Warn(format string, v ...interface{})

	// Error starts a new message with error level.
	//
	// You must call Msg on the returned event in order to send the event.
	Error(format string, v ...interface{})

	// Fatal starts a new message with fatal level. The os.Exit(1) function
	// is called by the Msg method, which terminates the program immediately.
	//
	// You must call Msg on the returned event in order to send the event.
	Fatal(format string, v ...interface{})

	// Printf sends a log event using debug level and no extra field.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...interface{})
}

type defaultLogger struct {
	level Level
}

func (m *defaultLogger) Level(lvl Level) Logger {
	return &defaultLogger{level: lvl}
}

func (m *defaultLogger) GetLevel() Level {
	return m.level
}

func (m *defaultLogger) Enabled() bool {
	return m.level != Disabled
}

func (m *defaultLogger) Discard() Logger {
	return &defaultLogger{level: Disabled}
}

func (m *defaultLogger) Trace(format string, v ...interface{}) {
	m.send(Trace, format, v...)
}

func (m *defaultLogger) Debug(format string, v ...interface{}) {
	m.send(Debug, format, v...)
}

func (m *defaultLogger) Info(format string, v ...interface{}) {
	m.send(Info, format, v...)
}

func (m *defaultLogger) Warn(format string, v ...interface{}) {
	m.send(Warning, format, v...)
}

func (m *defaultLogger) Error(format string, v ...interface{}) {
	m.send(Error, format, v...)
}

func (m *defaultLogger) Fatal(format string, v ...interface{}) {
	m.send(Fatal, format, v...)
}

func (m *defaultLogger) Printf(format string, v ...interface{}) {
	m.send(Debug, format, v...)
}

func (m defaultLogger) send(lev Level, format string, v ...interface{}) {
	if m.level > lev {
		return
	}
	switch lev {
	case Nop:
	case Disabled:
	default:
		log.Printf(format, v...)
	}
}

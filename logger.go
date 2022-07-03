package logz

// Logger represent structured logger abstraction layer
type Logger interface {

	// Level creates a child logger with the minimum accepted level set to level.
	Level(Level) Logger

	// With adds the field key with value to log message context.
	With(key string, v interface{}) Logger

	// GetLevel returns the current Level of Logger.
	GetLevel() Level

	// WithLevel starts a new Event with level. Unlike Fatal and Panic
	// methods, WithLevel does not terminate the program or stop the ordinary
	// flow of a goroutine when used with their respective levels.
	WithLevel(level Level) Event

	// Trace starts a new message with trace level.
	//
	// You must call Send on the returned event in order to newEvent the event.
	Trace() Event
	// Debug starts a new message with debug level.
	//
	// You must call Send on the returned event in order to newEvent the event.
	Debug() Event

	// Info starts a new message with info level.
	//
	// You must call Send on the returned event in order to newEvent the event.
	Info() Event

	// Warn starts a new message with warn level.
	//
	// You must callSend on the returned event in order to newEvent the event.
	Warn() Event

	// Error starts a new message with error level.
	//
	// You must call Send on the returned event in order to newEvent the event.
	Error() Event

	// Fatal starts a new message with fatal level. The os.Exit(1) function
	// is called by the Send method, which terminates the program immediately.
	Fatal() Event
}

type defLogger struct {
	level Level
}

func (m *defLogger) Level(l Level) Logger {
	return &defLogger{level: l}
}

func (m *defLogger) With(key string, v interface{}) Logger {
	return m
}

func (m *defLogger) GetLevel() Level {
	return m.level
}

func (m *defLogger) WithLevel(level Level) Event {
	return m.newEvent(level)
}

func (m *defLogger) Trace() Event {
	return m.newEvent(Trace)
}

func (m *defLogger) Debug() Event {
	return m.newEvent(Debug)
}

func (m *defLogger) Info() Event {
	return m.newEvent(Info)
}

func (m *defLogger) Warn() Event {
	return m.newEvent(Warning)
}

func (m *defLogger) Error() Event {
	return m.newEvent(Error)
}

func (m *defLogger) Fatal() Event {
	return m.newEvent(Fatal)
}

func (m defLogger) newEvent(lev Level) Event {
	return &defEvent{level: lev}
}

type defEvent struct {
	level Level
}

func (e *defEvent) Enabled() bool {
	return e.level != Disabled
}

func (e *defEvent) Discard() Event {
	e.level = Disabled
	return e
}

func (e *defEvent) With(key string, v interface{}) Event {
	return e
}

func (e *defEvent) Err(err error) Event {
	return e.With(ErrorStackFieldName, err.Error())
}

func (e *defEvent) Send(s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

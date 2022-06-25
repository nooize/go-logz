package logz

// Logger represent structured logger abstraction layer
type Logger interface {

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
}

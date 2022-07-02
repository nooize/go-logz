package logz

const (
	// Debug defines debug log level.
	Debug Level = iota
	// Info defines info log level.
	Info
	// Warning defines warn log level.
	Warning
	// Error defines error log level.
	Error
	// Fatal defines fatal log level.
	Fatal
	// Nop defines an absent log level.
	Nop
	// Disabled disables the logger.
	Disabled

	// Trace defines trace log level.
	Trace Level = -1
	// Values less than TraceLevel are handled as numbers.

	// LevelTrace is the value used for the trace level field.
	LevelTrace = "trace"
	// LevelDebug is the value used for the debug level field.
	LevelDebug = "debug"
	// LevelInfo is the value used for the info level field.
	LevelInfo = "info"
	// LevelWarn is the value used for the warn level field.
	LevelWarn = "warn"
	// LevelError is the value used for the error level field.
	LevelError = "error"
	// LevelFatalValue is the value used for the fatal level field.
	LevelFatal = "fatal"
)

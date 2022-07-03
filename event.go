package logz

// Logger represent structured logger abstraction layer
type Event interface {

	// Enabled return false if the Event is going to be filtered out by
	// log level or sampling.
	Enabled() bool

	// Discard disables the event so Send(f) won't print it.
	Discard() Event

	// With adds the field key with value to log message context.
	With(key string, v interface{}) Event

	// Err starts a new message with error level.
	Err(error) Event

	// Send fire a new message with defined in Logger level.
	Send(string, ...interface{})
}

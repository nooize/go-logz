package zero

import (
	"github.com/nooize/go-logz"
	"github.com/rs/zerolog"
)

// Logger represent structured logger abstraction layer
type zeroLogger struct {
	logger *zerolog.Logger
}

func (z *zeroLogger) Level(lvl logz.Level) logz.Logger {
	l := z.logger.Level(levelIn(lvl))
	z.logger = &l
	return z
}

// With adds the field key with value to log message context.
func (z *zeroLogger) With(key string, v interface{}) logz.Logger {
	c := z.logger.With()
	switch v.(type) {
	case int, int8, int16, int32, int64:
		c = c.Int(key, v.(int))
	}
	l := c.Logger()
	z.logger = &l
	return z
}

func (z *zeroLogger) GetLevel() logz.Level {
	return levelOut(z.logger.GetLevel())
}

func (z *zeroLogger) WithLevel(lvl logz.Level) logz.Event {
	return &zeroEvent{event: z.logger.WithLevel(levelIn(lvl))}
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Trace() logz.Event {
	return &zeroEvent{event: z.logger.Trace()}
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Debug() logz.Event {
	return &zeroEvent{event: z.logger.Debug()}
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Info() logz.Event {
	return &zeroEvent{event: z.logger.Info()}
}

// Warn starts a new message with warn level.
func (z *zeroLogger) Warn() logz.Event {
	return &zeroEvent{event: z.logger.Warn()}
}

// Error starts a new message with error level.
func (z *zeroLogger) Error() logz.Event {
	return &zeroEvent{event: z.logger.Error()}
}

// Err starts a new message with error level.
func (z *zeroLogger) Err(v error) logz.Event {
	return &zeroEvent{event: z.logger.Err(v)}
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
func (z *zeroLogger) Fatal() logz.Event {
	return &zeroEvent{event: z.logger.Fatal()}
}

package zero

import (
	"github.com/nooize/go-logz"
	"github.com/rs/zerolog"
)

// Logger represent structured logger abstraction layer
type zeroLogger struct {
	logger zerolog.Logger
}

func (z *zeroLogger) Level(lvl logz.Level) logz.Logger {
	return &zeroLogger{logger: z.logger.Level(levelIn(lvl))}
}

func (z *zeroLogger) GetLevel() logz.Level {
	return levelOut(z.logger.GetLevel())
}

func (z *zeroLogger) Enabled() bool {
	return z.Enabled()
}

func (z *zeroLogger) Discard() logz.Logger {
	z.logger = z.logger.Level(zerolog.Disabled)
	return z
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Trace(format string, v ...interface{}) {
	z.logger.Trace().Msgf(format, v...)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Debug(format string, v ...interface{}) {
	z.logger.Debug().Msgf(format, v...)
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Info(format string, v ...interface{}) {
	z.logger.Info().Msgf(format, v...)
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Warn(format string, v ...interface{}) {
	z.logger.Warn().Msgf(format, v...)
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Error(format string, v ...interface{}) {
	z.logger.Error().Msgf(format, v...)
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method, which terminates the program immediately.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroLogger) Fatal(format string, v ...interface{}) {
	z.logger.Fatal().Msgf(format, v...)
}

func (z *zeroLogger) Printf(format string, v ...interface{}) {
	z.logger.Printf(format, v...)
}

package zero

import (
	"github.com/nooize/go-logz"
	"github.com/rs/zerolog"
)

func levelOut(l zerolog.Level) logz.Level {
	switch l {
	case zerolog.TraceLevel:
		return logz.Trace
	case zerolog.DebugLevel:
		return logz.Debug
	case zerolog.InfoLevel:
		return logz.Info
	case zerolog.WarnLevel:
		return logz.Warning
	case zerolog.ErrorLevel:
		return logz.Error
	case zerolog.FatalLevel:
		return logz.Fatal
	case zerolog.NoLevel:
		return logz.Nop
	case zerolog.Disabled:
		return logz.Disabled
	}
	return logz.Nop
}

func levelIn(l logz.Level) zerolog.Level {
	switch l {
	case logz.Trace:
		return zerolog.TraceLevel
	case logz.Debug:
		return zerolog.DebugLevel
	case logz.Info:
		return zerolog.InfoLevel
	case logz.Warning:
		return zerolog.WarnLevel
	case logz.Error:
		return zerolog.ErrorLevel
	case logz.Fatal:
		return zerolog.FatalLevel
	case logz.Nop:
		return zerolog.NoLevel
	case logz.Disabled:
		return zerolog.Disabled
	}
	return zerolog.NoLevel
}

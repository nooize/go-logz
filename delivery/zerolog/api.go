package zero

import (
	"github.com/nooize/go-logz"
	"github.com/rs/zerolog"
	"io"
)

func New(w io.Writer, opts ...logz.LoggerOption) logz.Logger {
	out := logz.Logger(&zeroLogger{zerolog.New(w).With().Timestamp().Logger()})
	for _, opt := range opts {
		out = opt(out)
	}
	return out
}

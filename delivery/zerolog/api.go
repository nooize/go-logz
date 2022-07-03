package zero

import (
	"github.com/nooize/go-logz"
	"github.com/rs/zerolog"
	"io"
)

func New(w io.Writer, opts ...logz.LoggerOption) logz.Logger {
	zl := zerolog.New(w).With().Timestamp().Logger()
	out := logz.Logger(&zeroLogger{&zl})
	for _, opt := range opts {
		out = opt(out)
	}
	return out
}

package sys

import (
	"github.com/nooize/go-logz"
)

func New(l logz.Level) logz.Logger {
	return &sysLogger{
		level: l,
	}
}

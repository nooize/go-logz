package multi

import (
	"github.com/nooize/go-logz"
)

func New(logs ...logz.Logger) logz.Logger {
	out := multiLogger{
		active:  true,
		loggers: make([]logz.Logger, len(logs)),
	}
	for i, log := range logs {
		out.loggers[i] = log
	}
	return out
}

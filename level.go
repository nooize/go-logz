package logz

import (
	"strconv"
)

// Level defines log levels.
type Level int8

func (l Level) String() string {
	switch l {
	case Trace:
		return LevelTrace
	case Debug:
		return LevelDebug
	case Info:
		return LevelInfo
	case Warning:
		return LevelWarn
	case Error:
		return LevelError
	case Fatal:
		return LevelFatal
	case Disabled:
		return "disabled"
	case Nop:
		return ""
	}
	return strconv.Itoa(int(l))
}

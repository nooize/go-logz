package logz

import (
	"os"
	"sync"
)

var (
	rootLogger Logger
	loggerPool = make(map[string]Logger)
	loggersMu  sync.RWMutex
)

func Append(key string, log Logger) {
	if len(key) == 0 || log == nil {
		return
	}
	loggersMu.Lock()
	loggerPool[key] = log
	loggersMu.Unlock()
}

func Get(keys ...string) Logger {
	if len(keys) > 0 {
		for _, key := range keys {
			loggersMu.RLock()
			l, ok := loggerPool[key]
			loggersMu.RUnlock()
			if ok && l != nil {
				return l
			}
		}
	}
	return rootLogger
}

func init() {
	logLevel := Warning
	if os.Getenv("DEBUG_MODE") == "true" {
		logLevel = Debug
	}
	rootLogger = &defLogger{level: logLevel}
}

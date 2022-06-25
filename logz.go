package logz

import (
	"github.com/nooize/go-logz/delivery/sys"
	"os"
	"sync"
)

var (
	defaultLogger Logger
	loggerPool    = make(map[string]Logger)
	loggersMu     sync.RWMutex
)

func Append(key string, log Logger) {
	if len(key) == 0 || log == nil {
		return
	}
	loggersMu.Lock()
	loggerPool[key] = log
	loggersMu.Unlock()
}

func Log(keys ...string) (res Logger) {
	res = defaultLogger
	if len(keys) == 0 {
		return
	}
	loggersMu.RLock()
	if lc, ok := loggerPool[keys[0]]; ok && lc != nil {
		res = lc
	} else {
		//nl := res.Str("service", keys[0]).Logger()
		//res = &nl
	}
	loggersMu.RUnlock()
	return
}

func init() {
	logLevel := Warning
	if os.Getenv("DEBUG_MODE") == "true" {
		logLevel = Debug
	}
	defaultLogger = sys.New(logLevel)
}

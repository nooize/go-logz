package writers

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"time"
)

func NewRpcWriter(out io.Writer) io.Writer {
	cw := zerolog.ConsoleWriter{Out: out, NoColor: true, TimeFormat: time.RFC3339}
	cw.FormatLevel = emptyString
	cw.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	return cw
}

func emptyString(i interface{}) string {
	return ""
}

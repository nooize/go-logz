package writers

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"strings"
	"time"
)

func NewTextWriter(out io.Writer) io.Writer {
	cw := zerolog.ConsoleWriter{Out: out, TimeFormat: time.Stamp}
	cw.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	cw.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	cw.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	cw.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	return cw
}

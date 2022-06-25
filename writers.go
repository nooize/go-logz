package logz

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
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

func NewFileRotateWriter(path string) (io.Writer, error) {
	dir := filepath.Dir(path)
	if !isDirWritable(dir) {
		if err := os.MkdirAll(dir, 0744); err != nil {
			err = fmt.Errorf("can't create log directory: %s -> %s", dir, err.Error())
			defaultLogger.Error().Send(err.Error())
			return nil, err
		}
	}
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default
	}, nil
}

func isDirWritable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	if !info.IsDir() {
		return false
	}

	// Check if the user bit is enabled in file permission
	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false
	}

	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		return false
	}

	err = nil
	if uint32(os.Geteuid()) != stat.Uid {

		return false
	}

	return true
}

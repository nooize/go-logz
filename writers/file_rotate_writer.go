package writers

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

func NewFileRotateWriter(path string) (io.Writer, error) {
	dir := filepath.Dir(path)
	if !isDirWritable(dir) {
		if err := os.MkdirAll(dir, 0744); err != nil {
			return nil, fmt.Errorf("can't create log directory")
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

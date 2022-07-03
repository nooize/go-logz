package tools

import (
	"github.com/nooize/go-logz"
	zero "github.com/nooize/go-logz/delivery/zerolog"
	"github.com/nooize/go-logz/writers"
	"path/filepath"
)

func AppendRpcLogger(key string, path string) error {
	w, err := writers.NewFileRotateWriter(path)
	if err != nil {
		return err
	}
	zero.New(writers.NewRpcWriter(w),
		logz.WithLevel(logz.Debug),
		logz.AppendAs(key),
	)
	return nil
}

func AppendRpcLoggers(keys []string, dir string) error {
	for _, key := range keys {
		if err := AppendRpcLogger(key, filepath.Join(dir, key+".log")); err != nil {
			logz.Get().Error().Send("log open error : %v", err)
		}
	}
	return nil
}

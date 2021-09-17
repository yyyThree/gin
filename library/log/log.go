package log

import (
	"github.com/yyyThree/gin/config"
	"github.com/yyyThree/zap"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	if logger != nil {
		return logger
	}
	logger = zap.New(zap.Config{
		Env:    zap.Env(config.Config.App.Env),
		Writer: zap.Writer(config.Config.Log.Writer),
		LogDir: config.Config.Log.Dir,
	})
	return logger
}

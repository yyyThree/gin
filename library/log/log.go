package log

import (
	"gin/config"
	"gin/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.SugaredLogger

type log struct {
	logger *zap.SugaredLogger
}

func New() *log {
	return &log{
		logger: getLogger(),
	}
}

// 获取logger处理器
func getLogger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	// Log Level 哪种级别及以上的日志会被写入
	level := zapcore.InfoLevel
	if config.IsDev() {
		level = zapcore.DebugLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// AddStacktrace 哪种级别及以上的日志被写入堆栈信息
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)).Sugar()
	return logger
}

// 编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	zapConfig := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder, //将级别转换成大写
		TimeKey:     "time",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:     "caller",
		StacktraceKey: "stack",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1e6)
		},
	}
	return zapcore.NewJSONEncoder(zapConfig)
}

func (log *log) Debug(msg string, data ...constant.BaseMap) {
	if len(data) == 0 {
		log.logger.Debugw(msg)
	} else {
		log.logger.Debugw(msg, "data", data[0])
	}
}

func (log *log) Info(msg string, data ...constant.BaseMap) {
	if len(data) == 0 {
		log.logger.Infow(msg)
	} else {
		log.logger.Infow(msg, "data", data[0])
	}
}

func (log *log) Warn(msg string, data ...constant.BaseMap) {
	if len(data) == 0 {
		log.logger.Warnw(msg)
	} else {
		log.logger.Warnw(msg, "data", data[0])
	}
}

func (log *log) Error(msg string, data ...constant.BaseMap) {
	if len(data) == 0 {
		log.logger.Errorw(msg)
	} else {
		log.logger.Errorw(msg, "data", data[0])
	}
}

func (log *log) Panic(msg string, data ...constant.BaseMap) {
	if len(data) == 0 {
		log.logger.Panicw(msg)
	} else {
		log.logger.Panicw(msg, "data", data[0])
	}
}

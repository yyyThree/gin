package log

import (
	"gin/config"
	"gin/constant"
	"gin/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
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
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
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

// 日志输出
func getLogWriter() zapcore.WriteSyncer {
	// 日志切割设置，按 年/月/日 切割
	file := strings.TrimRight(config.Config.Log.Dir, "/") + "/" + helper.FormatDateNowBySlash() + ".log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,  // 日志文件位置
		MaxSize:    500,   // 日志文件最大大小(MB)
		MaxBackups: 2,     // 保留旧文件最大数量
		MaxAge:     0,     // 保留旧文件最长天数
		Compress:   false, // 是否压缩旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
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

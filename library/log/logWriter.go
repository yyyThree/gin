package log

import (
	"gin/config"
	"gin/helper"
	"gin/library/redis"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

// 为 logger 提供写入 redis 队列的 io 接口
type redisWriter struct {
}

// 日志输出
func getLogWriter() zapcore.WriteSyncer {
	switch config.Config.Log.Out {
	case "stdout":
		return getStdoutLogWriter()
	case "redis":
		return getRedisWriter()
	default: // 默认为文件存储
		return getFileLogWriter()
	}
}

// 屏幕输出
func getStdoutLogWriter() zapcore.WriteSyncer  {
	return zapcore.AddSync(os.Stdout)
}

// redis 实现 o.Writer 接口
func (r *redisWriter) Write(b []byte) (int, error) {
	// 日志切割设置，按 年/月/日 切割
	k := config.Config.Log.RedisKey + "." + helper.FormatDateNow()
	s := string(b)
	s = strings.TrimRight(s, "\n")
	n, err := redis.GetConn().RPush(redis.GetCtx(), k, s).Result()
	return int(n), err
}

// redis写入
func getRedisWriter() zapcore.WriteSyncer  {
	return zapcore.AddSync(&redisWriter{})
}

// 文件写入
func getFileLogWriter() zapcore.WriteSyncer  {
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

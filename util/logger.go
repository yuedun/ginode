package util

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://www.bilibili.com/read/cv18134162/
var SugarLogger *zap.SugaredLogger

func InitLogger() {
	// 写入位置
	writeSyncer := getLogWriter()
	// 编码格式
	encoder := getEncoder()
	// 需传入Encoder、WriterSyncer、Log Level
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer...), zapcore.DebugLevel)
	// 使用zap.New(…)方法来手动传递所有配置
	// 增加 Caller 信息
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

// 自定义日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// ISO8601TimeEncoder 序列化时间。以毫秒为精度的 ISO8601 格式字符串的时间。
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	// CapitalLevelEncoder 将Level序列化为全大写字符串。例如， InfoLevel被序列化为“INFO”。
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 在zap中加入Lumberjack支持
func getLogWriter() []zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/ginode.log",
		MaxSize:    128,   // 以 MB 为单位
		MaxBackups: 5,     // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxAge:     15,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return []zapcore.WriteSyncer{zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)}
}

// EncodeTime 自定义时间输出编码器
func EncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02-15:04:05.000"))
}

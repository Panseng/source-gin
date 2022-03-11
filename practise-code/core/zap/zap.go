package zap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"practise-code/config"
	"strings"
)

func InitSugLogger(cfg config.ZAP) *zap.SugaredLogger{
	writeSyncer := getLogWriter(cfg)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, getLevel(cfg.Level))
	return zap.New(core, zap.AddCaller()).Sugar()
}

// 获取日志打印层级
// 如果层级内容输入有误，则默认打印 zapcore.WarnLevel
func getLevel(l string) zapcore.Level {
	switch strings.ToLower(l) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.InfoLevel
	case "panic":
		return zapcore.WarnLevel
	case "fatal":
		return zapcore.ErrorLevel
	default: // 错配情况下，默认只输出警告以上内容
		return zapcore.WarnLevel
	}
}

// 使用Lumberjack进行日志切割归档
func getLogWriter(cfg config.ZAP) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Dir+cfg.Name,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAges,
		Compress:   cfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 写入日志的编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig) // JSON
	//return zapcore.NewConsoleEncoder(encoderConfig) // 行
}

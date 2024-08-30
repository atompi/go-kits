package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(logLevel string, logPath string, maxSize int, maxAge int, comp bool) *zap.Logger {
	writeSyncer := getLogWriter(logPath, maxSize, maxAge, comp)
	encoder := getEncoder()
	level := zapcore.InfoLevel
	switch logLevel {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logPath string, maxSize int, maxAge int, comp bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: logPath,
		MaxSize:  maxSize,
		MaxAge:   maxAge,
		Compress: comp,
	}
	return zapcore.AddSync(lumberJackLogger)
}

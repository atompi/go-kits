package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(logLevel string, logPath string, maxSize int, maxAge int, compress bool) *zap.Logger {
	debugWriteSyncer := getLogWriter(logPath+".debug.log", maxSize, maxAge, compress)
	infoWriteSyncer := getLogWriter(logPath+".info.log", maxSize, maxAge, compress)
	warnWriteSyncer := getLogWriter(logPath+".warn.log", maxSize, maxAge, compress)
	errorWriteSyncer := getLogWriter(logPath+".error.log", maxSize, maxAge, compress)

	encoder := getEncoder()

	debugCore := zapcore.NewCore(encoder, debugWriteSyncer, zapcore.DebugLevel)
	infoCore := zapcore.NewCore(encoder, infoWriteSyncer, zapcore.InfoLevel)
	warnCore := zapcore.NewCore(encoder, warnWriteSyncer, zapcore.WarnLevel)
	errorCore := zapcore.NewCore(encoder, errorWriteSyncer, zapcore.ErrorLevel)

	tee := zapcore.NewTee(debugCore, infoCore, warnCore, errorCore)

	logger := zap.New(tee)
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logPath string, maxSize int, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: logPath,
		MaxSize:  maxSize,
		MaxAge:   maxAge,
		Compress: compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

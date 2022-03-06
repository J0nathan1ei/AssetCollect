package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"manager/config"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.ManagerCfg.ZapCfg.Path,
		MaxSize:    config.ManagerCfg.ZapCfg.MaxSize,
		MaxBackups: config.ManagerCfg.ZapCfg.MaxBackups,
		MaxAge:     config.ManagerCfg.ZapCfg.MaxAge,
		Compress:   config.ManagerCfg.ZapCfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
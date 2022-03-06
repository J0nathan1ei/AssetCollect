package common

import (
	"agent/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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
		Filename:   config.AgentCfg.ZapCfg.Path,
		MaxSize:    config.AgentCfg.ZapCfg.MaxSize,
		MaxBackups: config.AgentCfg.ZapCfg.MaxBackups,
		MaxAge:     config.AgentCfg.ZapCfg.MaxAge,
		Compress:   config.AgentCfg.ZapCfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

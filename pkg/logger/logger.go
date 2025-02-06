package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.SugaredLogger {
    encoderCfg := zapcore.EncoderConfig{
        TimeKey:        "ts",
        MessageKey:     "msg",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        StacktraceKey:  "trace",
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.RFC3339TimeEncoder,
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
        LineEnding:     zapcore.DefaultLineEnding,
    }
    zlogger := zap.New(
        zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel),
        zap.AddCaller(),
        zap.AddStacktrace(zap.ErrorLevel),
    )
    return zlogger.Sugar()
}
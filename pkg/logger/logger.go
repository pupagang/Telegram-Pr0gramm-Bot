package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var InfoLogger *zap.Logger
var PanicLogger *zap.Logger
var ErrorLogger *zap.Logger

func init() {
	currentTime := time.Now().Format("2006-01-02")
	file, err := os.OpenFile(fmt.Sprintf("%s_logs.log", currentTime), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	InfoLogger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		file,
		zap.DebugLevel,
	))

	PanicLogger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		file,
		zap.PanicLevel,
	))

	ErrorLogger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		file,
		zap.ErrorLevel,
	))
}

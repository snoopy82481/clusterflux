package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func LogInit(d bool, f *os.File) {

	pe := zap.NewProductionEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)

	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level := zap.InfoLevel
	if d {
		level = zap.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	l := zap.New(core)

	log = l.Sugar()
}

func LogStart(functionName string) {
	log.Info("*********************************************\n" +
		"Started processing " + functionName + " at " + time.Now().Format(time.RFC3339) + "\n" +
		"*********************************************\n")
}

func LogStop(functionName string) {
	log.Info("*********************************************\n" +
		"Finished  processing " + functionName + " at " + time.Now().Format(time.RFC3339) + "\n" +
		"*********************************************\n")
}

func GetLogger() *zap.SugaredLogger {
	return log
}

func LogSuccess(message string) {
	log.Infow("SUCCESS",
		"timestamp", time.Now().Format(time.RFC3339),
		"message", message,
	)
}

func LogError(message string, err error) {
	log.Errorw("ERROR",
		"timestamp", time.Now().Format(time.RFC3339),
		"message", message,
		"err", err,
	)
}

func LogWarn(message string) {
	log.Warnw("WARN",
		"timestamp", time.Now().Format(time.RFC3339),
		"message", message,
	)
}

func LogInfo(message string, function string) {
	log.Infow("INFO",
		"timestamp", time.Now().Format(time.RFC3339),
		"function", function,
		"message", message,
	)
}

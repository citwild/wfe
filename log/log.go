package log

import "github.com/uber-go/zap"

var root = zap.New(zap.NewJSONEncoder())

func Root() zap.Logger {
	return root
}

func Debug(msg string, fields ...zap.Field) {
	root.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	root.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	root.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	root.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	root.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	root.Fatal(msg, fields...)
}

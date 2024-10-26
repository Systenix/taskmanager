package logger

import (
	"go.uber.org/zap"
)

var Log zap.Logger

func SetLogger(newLogger *zap.Logger) {
	Log = *newLogger
}

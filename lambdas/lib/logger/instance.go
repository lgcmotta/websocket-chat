package logger

import "go.uber.org/zap"

var Instance *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}

	Instance, _ = config.Build()
}

func Sync() {
	_ = Instance.Sync()
}

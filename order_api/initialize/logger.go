package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	// 使用zap日志
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

package main

import (
	"api/user_api/global"
	"api/user_api/initialize"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化配置信息
	initialize.InitConfig()
	// 初始化路由
	g := initialize.InitRouter()
	// 初始化验证器
	if err := initialize.InitTranslator("zh"); err != nil {
		panic(err)
	}
	// 初始化自定义验证器
	initialize.InitCustomizeValidator()

	port := global.Cfg.Port
	zap.L().Info("starting http server on port", zap.Int("port", port))
	err := g.Run(":" + strconv.Itoa(port))
	if err != nil {
		zap.L().Error("start http server failed", zap.Error(err))
		return
	}
}

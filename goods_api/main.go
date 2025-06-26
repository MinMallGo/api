package main

import (
	"api/goods_api/global"
	"api/goods_api/initialize"
	"api/goods_api/middleware"
	"api/goods_api/utils/register"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func main() {
	// 初始化gin
	g := gin.Default()
	// 使用中间件
	g.Use(middleware.CORS())

	// 初始化日志
	initialize.InitLogger()
	// 初始化配置信息
	initialize.InitConfig()
	// 初始化路由
	g = initialize.InitRouter(g)
	// 初始化验证器
	if err := initialize.InitTranslator("zh"); err != nil {
		panic(err)
	}
	// 初始化自定义验证器
	initialize.InitCustomizeValidator()
	// 初始化redis连接
	initialize.InitRedis()
	// 初始化grpc user-service 的连接
	initialize.InitGoosSrv()
	// 优雅地注册到注册中心
	rc := register.NewConsulRegistry(global.Cfg.Consul.Host, global.Cfg.Consul.Port)
	id := uuid.New().String()
	if err := rc.Register(&register.SrvRegisterArgs{
		Name: global.Cfg.Name,
		ID:   id,
		Host: global.Cfg.Host,
		Port: global.Cfg.Port,
		Tags: global.Cfg.Tags,
	}); err != nil {
		zap.L().Panic("register service failed", zap.Error(err))
	}
	// 注册结束

	port := global.Cfg.Port
	zap.L().Info("starting http server on port", zap.Int("port", port))
	go func() {
		err := g.Run(":" + strconv.Itoa(port))
		if err != nil {
			zap.L().Error("start http server failed", zap.Error(err))
			return
		}
	}()

	// 优雅地退出注册
	// 监听退出指令
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err := rc.Deregister(id)
	if err != nil {
		zap.L().Error("服务优雅退出失败", zap.Error(err))
	}
	zap.L().Info("服务已经优雅地退出", zap.String("id", id))
}

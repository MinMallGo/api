package initialize

import (
	"api/order_api/router"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(g *gin.Engine) *gin.Engine {
	// 添加一个健康检测的接口
	g.GET("health", func(context *gin.Context) {
		context.JSON(200, "")
		return
	})
	group := g.Group("/v1")
	// 注册购物车路由
	router.RegisterCartRouter(group)
	// 注册订单路由
	router.RegisterOrderRouter(group)
	return g
}

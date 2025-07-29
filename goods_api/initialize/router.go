package initialize

import (
	"goods_api/router"

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
	// 注册商品路由
	router.RegisterGoodsRouter(group)
	// 注册分类路由
	router.RegisterCategoryRouter(group)
	// 注册轮播图
	router.RegisterBannerRouter(group)
	// 注册品牌
	router.RegisterBrandRouter(group)
	// 注册品牌分类
	router.RegisterCategoryBrandRouter(group)
	return g
}

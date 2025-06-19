package router

import (
	"api/goods_api/api/goods"
	"api/goods_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGoodsRouter(group *gin.RouterGroup) {
	g := group.Group("/goods")
	{
		g.GET("list", goods.List)
		g.POST("create", middleware.JwtAuth(), middleware.AdminAuth(), goods.Create)
	}
}

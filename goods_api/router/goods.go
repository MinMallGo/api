package router

import (
	"goods_api/api/goods"
	"goods_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGoodsRouter(group *gin.RouterGroup) {
	g := group.Group("/goods")
	{
		g.GET("", goods.List)
		g.POST("", middleware.JwtAuth(), middleware.AdminAuth(), goods.Create)            // 创建商品
		g.DELETE("/:id", middleware.JwtAuth(), middleware.AdminAuth(), goods.Delete)      // 删除商品
		g.PUT("/:id", middleware.JwtAuth(), middleware.AdminAuth(), goods.Update)         // 更新商品信息
		g.PATCH("/:id", middleware.JwtAuth(), middleware.AdminAuth(), goods.UpdateStatus) // 仅仅更新状态
		g.GET("/:id/stock", goods.Stock)                                                  // 查看商品库存
	}
}

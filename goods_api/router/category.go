package router

import (
	"github.com/gin-gonic/gin"
	"goods_api/api/category"
	"goods_api/middleware"
)

func RegisterCategoryRouter(group *gin.RouterGroup) {
	g := group.Group("/category")
	{
		g.GET("", category.List)
		g.GET(":id", category.Detail)
		g.POST("", middleware.JwtAuth(), middleware.AdminAuth(), category.Create)       // 创建商品
		g.DELETE("/:id", middleware.JwtAuth(), middleware.AdminAuth(), category.Delete) // 删除商品
		g.PUT("/:id", middleware.JwtAuth(), middleware.AdminAuth(), category.Update)    // 更新商品信息
	}
}

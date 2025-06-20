package router

import (
	"api/goods_api/api/category"
	"api/goods_api/middleware"
	"github.com/gin-gonic/gin"
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

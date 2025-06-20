package router

import (
	bg "api/goods_api/api/category_brand"
	"api/goods_api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryBrandRouter(group *gin.RouterGroup) {
	g := group.Group("/category_brand")
	{
		g.GET("", bg.List)
		g.GET(":id", bg.Detail)
		g.POST("", middleware.JwtAuth(), middleware.AdminAuth(), bg.Create)
		g.DELETE("/:id", middleware.JwtAuth(), middleware.AdminAuth(), bg.Delete)
		g.PUT("/:id", middleware.JwtAuth(), middleware.AdminAuth(), bg.Update)
	}
}

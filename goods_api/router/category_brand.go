package router

import (
	"github.com/gin-gonic/gin"
	bg "goods_api/api/category_brand"
	"goods_api/middleware"
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

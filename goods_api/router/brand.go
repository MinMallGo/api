package router

import (
	"api/goods_api/api/brand"
	"api/goods_api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterBrandRouter(group *gin.RouterGroup) {
	g := group.Group("/brand")
	{
		g.GET("", brand.List)
		g.POST("", middleware.JwtAuth(), middleware.AdminAuth(), brand.Create)
		g.DELETE("/:id", middleware.JwtAuth(), middleware.AdminAuth(), brand.Delete)
		g.PUT("/:id", middleware.JwtAuth(), middleware.AdminAuth(), brand.Update)
	}
}

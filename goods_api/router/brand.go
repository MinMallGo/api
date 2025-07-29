package router

import (
	"github.com/gin-gonic/gin"
	"goods_api/api/brand"
	"goods_api/middleware"
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

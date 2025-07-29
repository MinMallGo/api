package router

import (
	"github.com/gin-gonic/gin"
	"goods_api/api/banner"
	"goods_api/middleware"
)

func RegisterBannerRouter(group *gin.RouterGroup) {
	g := group.Group("/banner")
	{
		g.GET("", banner.List)
		g.POST("", middleware.JwtAuth(), middleware.AdminAuth(), banner.Create)
		g.DELETE("/:id", middleware.JwtAuth(), middleware.AdminAuth(), banner.Delete)
		g.PUT("/:id", middleware.JwtAuth(), middleware.AdminAuth(), banner.Update)
	}
}

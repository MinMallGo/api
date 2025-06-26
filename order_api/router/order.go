package router

import (
	"api/order_api/api/order"
	"api/order_api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRouter(group *gin.RouterGroup) {
	g := group.Group("/order").Use(middleware.JwtAuth())
	{
		g.GET("", order.List)
		g.GET("/:id", order.Detail)
		g.POST("", order.Create)
	}
}

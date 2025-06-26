package router

import (
	"api/order_api/api/order"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRouter(group *gin.RouterGroup) {
	g := group.Group("/order")
	{
		g.GET("", order.List)
	}
}

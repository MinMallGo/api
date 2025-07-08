package router

import (
	"api/order_api/api/cart"
	"api/order_api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCartRouter(group *gin.RouterGroup) {
	g := group.Group("/cart").Use(middleware.JwtAuth()).Use(middleware.Trace())
	{
		g.GET("", cart.List)           // 获取购物车列表
		g.POST("", cart.Create)        // 添加到购物车
		g.DELETE("/:id", cart.Delete)  // 从购物车移除
		g.PUT("/", cart.Update)        // 更新购物车商品数量
		g.PATCH("/", cart.SelectGoods) // 勾选/批量勾选商品
	}
}

package router

import (
	"api/goods_api/api/goods"

	"github.com/gin-gonic/gin"
)

func RegisterGoodsRouter(group *gin.RouterGroup) {
	user := group.Group("/goods")
	{
		user.GET("list", goods.List)
	}
}

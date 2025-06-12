package initialize

import (
	"api/user_api/api"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	g := gin.Default()
	group := g.Group("/v1")
	{
		user := group.Group("/user")
		{
			user.GET("list", api.GetUserList)
			user.POST("pwd_login", api.PasswordLogin)
		}
	}
	return g
}

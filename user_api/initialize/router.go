package initialize

import (
	"api/user_api/router"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(g *gin.Engine) *gin.Engine {
	group := g.Group("/v1")
	// 注册用户路由
	router.RegisterUserRouter(group)
	// 注册基础路由
	router.RegisterUtilRouter(group)
	return g
}

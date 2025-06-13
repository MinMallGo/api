package router

import (
	"api/user_api/api"
	"api/user_api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(group *gin.RouterGroup) {
	user := group.Group("/user")
	{
		user.GET("list", middleware.JwtAuth(), middleware.AdminAuth(), api.GetUserList)
		user.POST("pwd_login", api.PasswordLogin)
		user.POST("register", api.UserCreate)
	}
}

package router

import (
	"api/user_api/api"
	"github.com/gin-gonic/gin"
)

// RegisterUtilRouter 注册一些常用的组件路由，比如说sms，图形验证码
func RegisterUtilRouter(group *gin.RouterGroup) {
	util := group.Group("/base")
	{
		util.GET("captcha", api.Captcha)
		util.GET("send_sms", api.SendSms)
	}
}

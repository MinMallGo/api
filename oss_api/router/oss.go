package router

import (
	"api/oss_api/api"
	"github.com/gin-gonic/gin"
)

func RegisterOSSRouter(group *gin.RouterGroup) {
	g := group.Group("/oss")
	{
		// 仅支持单文件上传，多文件上传请多次调用
		// 1. 能清楚的知道是否上传成功
		// 2. 避免大文件卡住
		// 3. 能更好维护
		g.POST("", api.Upload)
	}
}

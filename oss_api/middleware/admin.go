package middleware

import (
	"api/oss_api/structure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("claim")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "请先登录",
			})
			c.Abort()
			return
		}
		myClaim, ok := claim.(*structure.MyClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "参数异常",
			})
			c.Abort()
			return
		}

		if myClaim.AuthorizationId != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "无访问权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

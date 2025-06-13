package api

import (
	"api/user_api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Captcha(c *gin.Context) {
	id, captcha, err := utils.GenerateCaptcha()
	if err != nil {
		zap.L().Error("Captcha", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captcha_id": id,
		"captcha":    captcha,
	})
}

func SendSms(c *gin.Context) {
	//玩不了，这玩意儿个人没法认证了
}

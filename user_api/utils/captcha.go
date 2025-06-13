package utils

import (
	"api/user_api/global"
	"context"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha() (string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	if err != nil {
		return "", "", err
	}

	// id => answer
	_, err = global.Redis.Set(context.Background(), id, answer, 30*time.Minute).Result()
	if err != nil {
		zap.L().Error("Redis.Set", zap.Error(err))
		return "", "", err
	}

	return id, b64s, nil
}

func VerifyCaptcha(id, answer string) bool {
	// 读取并比较内容是否一致
	captcha, err := global.Redis.Get(context.Background(), id).Result()
	if err != nil {
		return false
	}
	global.Redis.Expire(context.Background(), id, 0)

	return captcha == answer
}

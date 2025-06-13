package test

import (
	"github.com/mojocn/base64Captcha"
	"log"
	"testing"
)

var store = base64Captcha.DefaultMemStore

// 创建一个简单的图形
func TestGenCode(t *testing.T) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	if err != nil {
		panic(err)
	}
	log.Println(id, b64s, answer)
}

package global

import (
	config2 "api/oss_api/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
)

var (
	Trans ut.Translator
	Cfg   *config2.Config = &config2.Config{}
	Redis *redis.Client   = &redis.Client{}
)

package initialize

import (
	"api/goods_api/global"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Cfg.Redis.Host, global.Cfg.Redis.Port),
		Password: global.Cfg.Redis.Password,
		DB:       0,
	})
}

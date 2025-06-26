package global

import (
	config2 "api/order_api/config"
	proto "api/order_api/proto/gen"

	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
)

type CrossService struct {
	OrderSrv  proto.OrderClient
	CartSrv   proto.CartClient
	Goods     proto.GoodsClient
	Inventory proto.InventoryClient // 库存服务，指的是 inventory
}

var (
	Trans    ut.Translator
	Cfg      *config2.Config = &config2.Config{}
	Redis    *redis.Client   = &redis.Client{}
	CrossSrv *CrossService   = &CrossService{}
	JWTUser  string          = "claim"
)

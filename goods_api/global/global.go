package global

import (
	config2 "api/goods_api/config"
	proto "api/goods_api/proto/gen"

	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
)

type GoodsService struct {
	Goods         proto.GoodsClient
	Brand         proto.BrandClient
	Banner        proto.BannerClient
	Category      proto.CategoryClient
	CategoryBrand proto.CategoryBrandClient
}

var (
	Trans    ut.Translator
	Cfg      *config2.Config = &config2.Config{}
	Redis    *redis.Client   = &redis.Client{}
	GoodsSrv *GoodsService   = &GoodsService{}
)

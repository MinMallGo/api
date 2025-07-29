package initialize

import (
	"fmt"
	"goods_api/global"
	"goods_api/middleware"
	proto "goods_api/proto/gen"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGoosSrv() {
	conn, err := grpc.NewClient(
		fmt.Sprintf(`consul://%s:%d/%s?wait=14s`, global.Cfg.Consul.Host, global.Cfg.Consul.Port, global.Cfg.GoodsServer.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(middleware.UnaryServerInterceptor()), // 添加一个限流器
	)
	if err != nil {
		zap.L().Fatal("【InitGoosSrv】服务获取失败：", zap.Error(err))
	}
	zap.L().Info("<UNK>InitGoosSrv<UNK>", zap.Any("<conn_consul>", fmt.Sprintf(`consul://%s:%d/%s?wait=14s`, global.Cfg.Consul.Host, global.Cfg.Consul.Port, global.Cfg.GoodsServer.Name)))

	global.GoodsSrv = &global.GoodsService{
		Goods:         proto.NewGoodsClient(conn),
		Brand:         proto.NewBrandClient(conn),
		Banner:        proto.NewBannerClient(conn),
		Category:      proto.NewCategoryClient(conn),
		CategoryBrand: proto.NewCategoryBrandClient(conn),
	}
}

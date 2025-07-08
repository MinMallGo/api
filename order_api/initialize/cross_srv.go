package initialize

import (
	"api/order_api/global"
	proto "api/order_api/proto/gen"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitCrossSrv() {
	ConnOrder()
	ConnGoods()
	ConnStock()
}

// ConnOrder 连接订单服务
func ConnOrder() {
	conn, err := grpc.NewClient(
		fmt.Sprintf(`consul://%s:%d/%s?wait=14s`, global.Cfg.Consul.Host, global.Cfg.Consul.Port, global.Cfg.CrossSrv.OrderSrv),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("【%s】服务获取失败：", global.Cfg.CrossSrv.OrderSrv), zap.Error(err))
	}

	global.CrossSrv.OrderSrv = proto.NewOrderClient(conn)
	global.CrossSrv.CartSrv = proto.NewCartClient(conn)
}

// ConnGoods goods-service
func ConnGoods() {
	conn, err := grpc.NewClient(
		fmt.Sprintf(`consul://%s:%d/%s?wait=14s`, global.Cfg.Consul.Host, global.Cfg.Consul.Port, global.Cfg.CrossSrv.GoodsSrv),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("【%s】服务获取失败：", global.Cfg.CrossSrv.GoodsSrv), zap.Error(err))
	}

	global.CrossSrv.Goods = proto.NewGoodsClient(conn)
}

// ConnStock stock-service
func ConnStock() {
	conn, err := grpc.NewClient(
		fmt.Sprintf(`consul://%s:%d/%s?wait=14s`, global.Cfg.Consul.Host, global.Cfg.Consul.Port, global.Cfg.CrossSrv.InventorySrv),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("【%s】服务获取失败：", global.Cfg.CrossSrv.InventorySrv), zap.Error(err))
	}

	global.CrossSrv.Inventory = proto.NewInventoryClient(conn)
}

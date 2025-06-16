package initialize

import (
	"api/user_api/global"
	"api/user_api/proto"
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUserSrv() {
	conn, err := grpc.NewClient(
		fmt.Sprintf(`consul://%s:%d/%s?wait=14s`, global.Cfg.Consul.Host, global.Cfg.Consul.Port, global.Cfg.UserServer.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.L().Fatal("【InitUserSrv】服务获取失败：", zap.Error(err))
	}

	global.UserSrv = proto.NewUserClient(conn)
}

func InitUserSrvDeprecated() {
	/*
		1. 通过配置文件获取consul的连接信息
		2. 通过服务发现获取服务
		3. 通过服务的ip:port连接grpc服务
		4. 保存连接
	*/
	srv, err := filterService(global.Cfg.UserServer.Name)
	if err != nil {
		zap.L().Fatal("[InitUserSrv] 连接consul获取服务失败", zap.Error(err))
	}

	host, port := "", 0
	for _, service := range srv {
		host = service.Address
		port = service.Port
	}

	client, err := grpc.NewClient(fmt.Sprintf("%v:%v", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Fatal("[InitUserSrv] grpc-user-srv 连接失败", zap.Error(err))
	}
	global.UserSrv = proto.NewUserClient(client)
}

func filterService(name string) (map[string]*api.AgentService, error) {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", global.Cfg.Consul.Host, global.Cfg.Consul.Port)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
}

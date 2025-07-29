package initialize

import (
	"go.uber.org/zap"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		zap.S().Fatal(err)
		return
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "grpc:/Goods/GoodsList",
			Threshold:              2,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       10000,
		},
	})

	if err != nil {
		zap.S().Fatal(err)
		return
	}
}

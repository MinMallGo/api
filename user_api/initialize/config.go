package initialize

import (
	"api/user_api/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
)

/*
	1. 从环境变量里读取是否是debug环境
	2. 根据是否debug环境读取不同的配置文件
	3. 解析文件到结构体
	4. 将读取到的配置文件赋予全局config变量
*/

func GetEnv(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	curPath, _ := os.Getwd()
	configName := "config.yaml"
	debug := GetEnv("GoProjectDebug")
	if debug {
		configName = "config_debug.yaml"
	}

	v := viper.New()
	// TODO 监听文件变化 以及替换项目中的东西

	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath(curPath)
	err := v.ReadInConfig()
	if err != nil {
		zap.L().Info("读取配置文件错误", zap.Error(err))
		return
	}

	err = v.Unmarshal(global.Cfg)
	if err != nil {
		zap.L().Info("读取配置文件错误", zap.Error(err))
		return
	}
	log.Println(global.Cfg)

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.Cfg)
		zap.L().Info("配置文件发生了变化", zap.Any("config", global.Cfg))
	})
	v.WatchConfig()

}

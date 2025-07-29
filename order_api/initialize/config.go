package initialize

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"api/order_api/config"
	"api/order_api/global"
)

/*
	1. 从环境变量里读取是否是debug环境
	2. 根据是否debug环境读取不同的配置文件
	3. 解析文件到结构体
	4. 将读取到的配置文件赋予全局config变量
*/

var (
	nacos = &config.NacosCnf{}
)

func GetEnv(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// 修改成从本地读取nacos的配置文件
// 然后从nacos获取配置，再然后再解析nacos的配置文件到本地

func InitConfig() {
	curPath, _ := os.Getwd()
	configName := "config.json"
	debug := GetEnv("GoProjectDebug")
	if debug {
		configName = "config_debug.json"
	}
	//curPath = `C:\ezgo\api\order_api`
	file, err := os.ReadFile(filepath.Join(curPath, configName))
	if err != nil {
		zap.L().Fatal("[InitConfig] 读取nacos配置文件失败:", zap.Error(err))
	}

	err = json.Unmarshal(file, nacos)
	if err != nil {
		zap.L().Fatal("[InitConfig] 解析nacos配置文件失败:", zap.Error(err))
	}

	// 读取nacos的配置文件之后，打印以下看看
	fmt.Printf("nacos配置文件：%#v\n", nacos)

	// 连接nacos获取配置
	cnf, err := getConfig()
	if err != nil {
		zap.L().Fatal("[InitConfig].[getConfig] with error:", zap.Error(err))
	}

	err = json.Unmarshal([]byte(cnf), &global.Cfg)
	if err != nil {
		zap.L().Fatal("[InitConfig].[json.Unmarshal] with error:", zap.Error(err))
	}
	zap.L().Info("[InitConfig] <UNK>", zap.Any("cnf", cnf))
}

func getConfig() (string, error) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacos.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            nacos.Username, // Nacos服务端的API鉴权Username
		Password:            nacos.Password, // Nacos服务端的API鉴权Password
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacos.Host,
			ContextPath: "/nacos",
			Port:        nacos.Port,
			Scheme:      "http",
		},
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		return "", err
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacos.DataID,
		Group:  nacos.Group,
	})
	if err != nil {
		return "", err
	}

	return content, nil
}

func InitConfig2() {
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

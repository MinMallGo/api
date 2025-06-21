package config

type Config struct {
	Name   string    `mapstructure:"name" json:"name"`
	Host   string    `mapstructure:"host" json:"host"`
	Port   int       `mapstructure:"port" json:"port"`
	Tags   []string  `mapstructure:"tags" json:"tags"`
	Jwt    JwtSrv    `mapstructure:"jwt" json:"jwt"`
	Redis  RedisCnf  `mapstructure:"redis" json:"redis"`
	Consul ConsulCnf `mapstructure:"consul" json:"consul"`
	Oss    OSSCnf    `mapstructure:"oss" json:"oss"`
}

type JwtSrv struct {
	Key string `mapstructure:"key" json:"key"`
}

type RedisCnf struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulCnf struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosCnf struct {
	Host      string `json:"host"`
	Port      uint64 `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Namespace string `json:"namespace"`
	DataID    string `json:"data_id"`
	Group     string `json:"group"`
}

type OSSCnf struct {
	AccessKey string `mapstructure:"access_key" json:"access_key"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	Bucket    string `mapstructure:"bucket" json:"bucket"`
	Endpoint  string `mapstructure:"endpoint" json:"endpoint"`
}

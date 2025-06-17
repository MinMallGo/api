package config

type Config struct {
	Name       string    `mapstructure:"name" json:"name"`
	Port       int       `mapstructure:"port" json:"port"`
	UserServer UserSrv   `mapstructure:"user_srv" json:"user_srv"`
	Jwt        JwtSrv    `mapstructure:"jwt" json:"jwt"`
	Redis      RedisCnf  `mapstructure:"redis" json:"redis"`
	Consul     ConsulCnf `mapstructure:"consul" json:"consul"`
}

type UserSrv struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
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

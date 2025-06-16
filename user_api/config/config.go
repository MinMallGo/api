package config

type Config struct {
	Name       string    `mapstructure:"name"`
	Port       int       `mapstructure:"port"`
	UserServer UserSrv   `mapstructure:"user_srv"`
	Jwt        JwtSrv    `mapstructure:"jwt"`
	Redis      RedisCnf  `mapstructure:"redis"`
	Consul     ConsulCnf `mapstructure:"consul"`
}

type UserSrv struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type JwtSrv struct {
	Key string `mapstructure:"key"`
}

type RedisCnf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type ConsulCnf struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

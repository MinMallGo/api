package config

type Config struct {
	Name       string   `mapstructure:"name"`
	Port       int      `mapstructure:"port"`
	UserServer UserSrv  `mapstructure:"user_srv"`
	Jwt        JwtSrv   `mapstructure:"jwt"`
	Redis      RedisCnf `mapstructure:"redis"`
}

type UserSrv struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JwtSrv struct {
	Key string `mapstructure:"key"`
}

type RedisCnf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

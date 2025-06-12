package config

type Config struct {
	Name       string  `mapstructure:"name"`
	Port       int     `mapstructure:"port"`
	UserServer UserSrv `mapstructure:"user_srv"`
}

type UserSrv struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

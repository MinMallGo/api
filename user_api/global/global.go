package global

import (
	config2 "api/user_api/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans ut.Translator
	Cfg   *config2.Config = &config2.Config{}
)

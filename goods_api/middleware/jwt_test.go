package middleware

import (
	config2 "goods_api/config"
	"goods_api/global"
	"goods_api/structure"
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestToken(t *testing.T) {
	global.Cfg = &config2.Config{
		Jwt: config2.JwtSrv{
			Key: "3JAKLsijrLit0kK0oqyUBjCkUDyDdlwz1D0mXjAtgNc=",
		},
	}
	j := NewJwt()
	token, err := j.CreateToken(structure.MyClaims{
		ID:              11,
		NickName:        "ddff",
		AuthorizationId: 2,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "xx",
			Subject:   "xx",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  nil,
			ID:        "xx",
		},
	})
	if err != nil {
		log.Panicln(err)
	}
	log.Println(token)

	parse, err := j.JwtParse(token)
	if err != nil {
		panic(err)
	}
	log.Println(parse)
}

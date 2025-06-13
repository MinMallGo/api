package middleware

import (
	config2 "api/user_api/config"
	"api/user_api/global"
	"api/user_api/structure"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	global.Cfg = &config2.Config{
		Jwt: config2.JwtSrv{
			Key: "3JAKLsijrLit0kK0oqyUBjCkUDyDdlwz1D0mXjAtgNc=",
		},
	}
	j := NewJwt()
	token, err := j.CreateToken(structure.MyClaims{
		ID:              "1",
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

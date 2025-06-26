package structure

import "github.com/golang-jwt/jwt/v5"

type MyClaims struct {
	ID              int    `json:"id"`
	NickName        string `json:"nickname"`
	AuthorizationId int    `json:"authorizationId"`
	jwt.RegisteredClaims
}

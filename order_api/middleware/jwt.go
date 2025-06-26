package middleware

import (
	"api/order_api/global"
	"api/order_api/structure"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "请登录",
			})
			c.Abort()
			return
		}
		j := NewJwt()
		claim, err := j.JwtParse(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set(global.JWTUser, claim)
		c.Next()
	}
}

type Jwt struct {
	signingKey []byte
}

func NewJwt() *Jwt {
	return &Jwt{
		signingKey: []byte(global.Cfg.Jwt.Key),
	}
}

func (j *Jwt) CreateToken(claims structure.MyClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(j.signingKey)
}

func (j *Jwt) JwtParse(tokenString string) (*structure.MyClaims, error) {
	claims := &structure.MyClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return j.signingKey, nil })
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func (j *Jwt) JwtRenewal() {}

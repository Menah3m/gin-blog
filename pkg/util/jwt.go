package util

import (
	"github.com/Menah3m/gin-blog/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	username string `json:"username"`
	password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(username, password string) (string, error)  {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)  // 失效时间设置为3小时

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error)  {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid{
			return claims, nil
		}
	}

	return nil, err
}

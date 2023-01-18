package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("关注永雏塔菲喵")

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	claims := CustomClaims{
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "永雏塔菲", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

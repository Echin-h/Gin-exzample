package jwt

import (
	"LearningGo/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 把claims作为一个结构体
type Payload struct {
	Authorized bool   `json:"authorized"`
	User       string `json:"user"`
}

type MyCustomClaims struct {
	Payload
	jwt.RegisteredClaims
}

func NewToken(name string) (string, error) {

	//设置一些预定义  Payload
	claims := &MyCustomClaims{
		Payload: Payload{
			Authorized: true,
			User:       name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    configs.JwtSettings.Issuer,
			Subject:   configs.JwtSettings.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			// jwt.NewNumericDate 可以创建一个符合JWT标准的时间格式
		},
	}

	// 创建一个新的令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Header,token是一个对象

	//签名并获取完整的编码令牌作为字符串  Signature
	tokenString, err := token.SignedString([]byte(configs.JwtSettings.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

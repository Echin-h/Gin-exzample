package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// ParseToken 是 解析令牌
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	claims := &MyCustomClaims{}
	// 是*token和string之间的转换
	// 这是一个回调函数具体结构就是 jwt.Parse(string,KeyFunc)
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法 HMAC-SHA56签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return claims, nil
}

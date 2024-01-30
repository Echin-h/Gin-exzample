package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

// ParseToken 是 解析令牌
func ParseToken(bearerToken string) (*MyCustomClaims, error) {
	// 解析方式需要添加 Bearer token模式
	tokenParts := strings.Split(bearerToken, " ")                           //通过空格分隔出两个部分，并且存入数组之中
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" { // strings.ToLower会把字符串变成小写的形式
		return nil, errors.New("Invalid token format,you need add bearer")
	}
	tokenString := tokenParts[1]
	// 解析后续token
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

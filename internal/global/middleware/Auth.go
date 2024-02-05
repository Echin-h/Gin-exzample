package middleware

import (
	"LearningGo/internal/global/jwt"
	"LearningGo/internal/global/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to fetch token",
			})
			return
		}
		parseToken, err := jwt.ParseToken(token)
		if err != nil {
			log.SugarLogger.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":   "Failed to Auth",
				"error": err,
			})
			c.Abort()
			return
		}
		c.Set("Payload", parseToken)
		c.Next()
	}
}

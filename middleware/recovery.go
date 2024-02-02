package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecoveryWithWriter(os.Stderr, func(c *gin.Context, err interface{}) {

	})
}

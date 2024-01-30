package handler

import (
	"LearningGo/db"
	"LearningGo/jwt"
	"LearningGo/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Read(c *gin.Context) {
	payload, exists := c.Get("Payload")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": errors.New("Failed to Auth"),
		})
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	var user model.User
	db.DB.Where("name = ?", load.User).First(&user)
	c.JSON(http.StatusOK, gin.H{
		"it is:": "你的用户信息",
		"msg":    user,
	})
}

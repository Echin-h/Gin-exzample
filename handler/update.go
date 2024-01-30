package handler

import (
	"LearningGo/db"
	"LearningGo/jwt"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 修改密码

func Update(c *gin.Context) {
	// 输入老的密码和新的密码
	NewPassword := c.PostForm("NewPassword")
	payload, exists := c.Get("Payload")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "账号未登录",
		})
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	db.DB.Model(&model.User{}).Where("name = ? ", load.User).Update("password", NewPassword)
	c.JSON(http.StatusOK, gin.H{
		"msg":         "修改密码成功",
		"user":        load.User,
		"NewPassword": NewPassword,
	})
}

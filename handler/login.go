package handler

import (
	"LearningGo/db"
	"LearningGo/jwt"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户的登录
func Login(c *gin.Context) {
	// 用PostForm 传输数据
	name := c.PostForm("name")
	password := c.PostForm("password")

	var v2 model.User
	// 判断用户是否存在
	if tx := db.DB.Where(" name = ? ", name).First(&v2); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户不存在",
			"err": tx.Error,
		})
		return
	}

	// 判断密码是否正确
	if password != v2.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "密码不正确，小笨蛋",
		})
		return
	} else {
		token, err := jwt.NewToken(name)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":   "生成token失败",
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功，大聪明",
			"user":  v2,
			"token": token,
		})
	}
}

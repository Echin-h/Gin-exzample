package handler

import (
	"LearningGo/db"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code": 400,
			"msg":  err.Error(),
		})
		return
	}
	//判断姓名，密码，邮箱不能为空
	if len(user.Name) == 0 || len(user.Password) == 0 || len(user.Email) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Username, password or email is required...",
		})
		return
	}
	// 注册逻辑，是否用户名和邮箱已经存在
	var v1 model.User
	result := db.DB.Where("name = ?", user.Name).First(&v1)
	if result.Error == nil {
		c.JSON(200, gin.H{
			"msg": "姓名已经重复了",
			"err": result.Error,
		})
	} else {
		//保存user,如果已经存在了user表，就会直接插入数据
		if err := db.DB.Create(&user).Error; err != nil {
			err1 := db.DB.AutoMigrate(&user)
			if err1 != nil {
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

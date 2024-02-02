package handler

import (
	"LearningGo/db"
	"LearningGo/errs"
	"LearningGo/jwt"
	"LearningGo/log"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
)

// 用户的登录
func Login(c *gin.Context) {
	// 用PostForm 传输数据
	name := c.PostForm("name")
	password := c.PostForm("password")

	var v2 model.User
	if tx := db.DB.Where(" name = ? ", name).First(&v2); tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(tx.Error))
		return
	}

	if password != v2.Password {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("密码错误"))
		return
	}

	token, err := jwt.NewToken(name)
	if err != nil {
		log.SugarLogger.Error(err)
		errs.Fail(c, errs.UNTHORIZATION.WithOrigin(err))
		return
	}

	errs.Success(c, v2, map[string]string{"token": token})
}

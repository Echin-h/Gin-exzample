package user

import (
	"LearningGo/internal/global/db"
	errs2 "LearningGo/internal/global/errs"
	"LearningGo/internal/global/jwt"
	"LearningGo/internal/global/log"
	"LearningGo/internal/model"
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
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(tx.Error))
		return
	}

	if password != v2.Password {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("密码错误"))
		return
	}

	token, err := jwt.NewToken(name)
	if err != nil {
		log.SugarLogger.Error(err)
		errs2.Fail(c, errs2.UNTHORIZATION.WithOrigin(err))
		return
	}

	errs2.Success(c, v2, map[string]string{"token": token})
}

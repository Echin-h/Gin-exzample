package handler

import (
	"LearningGo/db"
	"LearningGo/errs"
	"LearningGo/log"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.SugarLogger.Error(err)
		errs.Fail(c, errs.INVALID_REQUEST.WithOrigin(err))
		return
	}

	if len(user.Name) == 0 || len(user.Password) == 0 || len(user.Email) == 0 {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("Username, password or email is required..."))
		return
	}

	var v1 model.User
	result := db.DB.Where("name = ?", user.Name).First(&v1)
	if result.Error == nil {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("姓名重复"))
	}

	if err := db.DB.Create(&user).Error; err != nil {
		err1 := db.DB.AutoMigrate(&user)
		if err1 != nil {
			log.SugarLogger.Error(err1)
			return
		}
		log.SugarLogger.Error(err)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(err))
		return
	}

	errs.Success(c, "注册成功")
}

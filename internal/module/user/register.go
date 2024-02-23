package user

import (
	"LearningGo/internal/global/casbin"
	"LearningGo/internal/global/db"
	errs2 "LearningGo/internal/global/errs"
	"LearningGo/internal/global/log"
	"LearningGo/internal/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.SugarLogger.Error(err)
		errs2.Fail(c, errs2.INVALID_REQUEST.WithOrigin(err))
		return
	}

	if len(user.Name) == 0 || len(user.Password) == 0 || len(user.Email) == 0 {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("Username, password or email is required..."))
		return
	}

	var v1 model.User
	result := db.DB.Where("name = ?", user.Name).First(&v1)
	if result.Error == nil {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("姓名重复"))
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.SugarLogger.Error(err)
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(err))
		return
	}
	err := casbin.Enforce.LinkUserWithPolicy(user.Name)
	if err != nil {
		log.SugarLogger.Error(err)
		errs2.Fail(c, errs2.SERVE_INTERNAL.WithOrigin(err))
		return
	}
	errs2.Success(c, "注册成功")
}

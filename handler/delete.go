package handler

import (
	"LearningGo/db"
	"LearningGo/errs"
	"LearningGo/jwt"
	"LearningGo/log"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
)

// 注销账户
func Delete(c *gin.Context) {
	password := c.PostForm("password")
	payload, exists := c.Get("payload")
	if !exists {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("没有获取到payload"))
	}

	load := payload.(*jwt.MyCustomClaims)

	var user model.User
	db.DB.Where("name = ?", load.User).First(&user)
	if user.Password != password {
		errs.Fail(c, errs.INVALID_REQUEST.WithTips("输入密码错误"))
		return
	}

	result := db.DB.Delete(&user)
	if result.Error != nil {
		log.SugarLogger.Error(result.Error)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(result.Error))
		return
	}

	errs.Success(c, "注销成功")
}

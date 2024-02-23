package user

import (
	"LearningGo/internal/global/casbin"
	"LearningGo/internal/global/db"
	errs2 "LearningGo/internal/global/errs"
	"LearningGo/internal/global/jwt"
	"LearningGo/internal/global/log"
	"LearningGo/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 注销账户
func Delete(c *gin.Context) {
	password := c.PostForm("password")
	payload, exists := c.Get("Payload")
	if !exists {
		errs2.Fail(c, errs2.UNTHORIZATION.WithTips("没有获取到payload"))
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	ok := casbin.Enforce.CheckUserPolicyForRead(load.User, "users", "write")
	if !ok {
		errs2.Fail(c, errs2.UNTHORIZATION.WithTips("没有权限修改"))
		return
	}
	var user model.User
	db.DB.Where("name = ?", load.User).First(&user)
	if user.Password != password {
		fmt.Println("user.Password", user.Password)
		fmt.Println(password)
		errs2.Fail(c, errs2.INVALID_REQUEST.WithTips(password))
		return
	}

	result := db.DB.Delete(&user)
	if result.Error != nil {
		log.SugarLogger.Error(result.Error)
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(result.Error))
		return
	}

	errs2.Success(c, "注销成功")
}

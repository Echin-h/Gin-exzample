package handler

import (
	"LearningGo/db"
	"LearningGo/errs"
	"LearningGo/jwt"
	"LearningGo/log"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
)

// 修改密码

func Update(c *gin.Context) {
	NewPassword := c.PostForm("NewPassword")
	payload, exists := c.Get("Payload")
	if !exists {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("无法get你的payload"))
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	tx := db.DB.Model(&model.User{}).Where("name = ? ", load.User).Update("password", NewPassword)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
	}
	errs.Success(c, load.User, NewPassword, "请妥善保存你的密码")
}

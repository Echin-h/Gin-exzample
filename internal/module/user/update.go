package user

import (
	"LearningGo/internal/global/db"
	errs2 "LearningGo/internal/global/errs"
	"LearningGo/internal/global/jwt"
	"LearningGo/internal/global/log"
	"LearningGo/internal/model"
	"github.com/gin-gonic/gin"
)

// 修改密码

func Update(c *gin.Context) {
	NewPassword := c.PostForm("NewPassword")
	payload, exists := c.Get("Payload")
	if !exists {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("无法get你的payload"))
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	tx := db.DB.Model(&model.User{}).Where("name = ? ", load.User).Update("password", NewPassword)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
	}
	errs2.Success(c, load.User, NewPassword, "请妥善保存你的密码")
}

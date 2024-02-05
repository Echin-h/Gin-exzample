package user

import (
	"LearningGo/internal/global/db"
	errs2 "LearningGo/internal/global/errs"
	"LearningGo/internal/global/jwt"
	"LearningGo/internal/global/log"
	"LearningGo/internal/model"
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	payload, exists := c.Get("Payload")
	if !exists {
		errs2.Fail(c, errs2.UNTHORIZATION.WithTips("Failed to Auth"))
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	var user model.User
	tx := db.DB.Where("name = ?", load.User).First(&user)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(tx.Error))
	}
	errs2.Success(c, user)
}

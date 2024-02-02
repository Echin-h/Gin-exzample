package handler

import (
	"LearningGo/db"
	"LearningGo/errs"
	"LearningGo/jwt"
	"LearningGo/log"
	"LearningGo/model"
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	payload, exists := c.Get("Payload")
	if !exists {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("Failed to Auth"))
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	var user model.User
	tx := db.DB.Where("name = ?", load.User).First(&user)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(tx.Error))
	}
	errs.Success(c, user)
}

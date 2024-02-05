package user

import (
	"LearningGo/internal/global/middleware"
	"github.com/gin-gonic/gin"
)

// user 主要用于封装路由列表

func (u *ModuleUser) InitRouter(r *gin.RouterGroup) {
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.PUT("/update", middleware.Auth(), Update)
	r.GET("/read", middleware.Auth(), Read)
	r.DELETE("/delete", middleware.Auth(), Delete)
}

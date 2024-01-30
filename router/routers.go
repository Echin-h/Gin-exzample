package router

import (
	"LearningGo/handler"
	"LearningGo/middleware"
	"github.com/gin-gonic/gin"
)

// router 主要用于封装路由列表

func PathUser(r *gin.Engine) {
	rootPath := r.Group("/user")
	{
		rootPath.POST("/register", handler.Register)
		rootPath.POST("/login", handler.Login)
		rootPath.PUT("/update", middleware.Auth(), handler.Update)
		//rootPath.GET("/read", middleware.Auth(), handler.Read)
	}
}

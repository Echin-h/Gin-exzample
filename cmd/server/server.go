package server

import (
	"LearningGo/configs"
	"LearningGo/internal/global/casbin"
	"LearningGo/internal/global/db"
	"LearningGo/internal/global/log"
	"LearningGo/internal/global/middleware"
	"LearningGo/internal/module"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	configs.Init()
	db.Init()
	casbin.Init()
	for _, m := range module.Modules {
		fmt.Println("Init Module: " + m.GetName())
		m.Init()
	}
}

func Run() {
	r := gin.New()
	r.Use(log.Init(), middleware.Recovery())

	for _, m := range module.Modules {
		fmt.Println("InitRouter: " + m.GetName())
		m.InitRouter(r.Group("/" + m.GetName()))
	}

	panic(r.Run(":9090"))
}

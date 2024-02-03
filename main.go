package main

import (
	"LearningGo/configs"
	"LearningGo/db"
	"LearningGo/log"
	"LearningGo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	log.InitLogger()
	configs.Init()
	db.Init()
	router.PathUser(r)

	panic(r.Run(":9090"))
}

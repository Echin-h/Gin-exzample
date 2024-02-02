package main

import (
	"LearningGo/db"
	"LearningGo/log"
	"LearningGo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()
	log.InitLogger()
	router.PathUser(r)

	panic(r.Run(":9090"))
}

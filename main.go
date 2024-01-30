package main

import (
	"LearningGo/db"
	"LearningGo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()
	router.PathUser(r)

	panic(r.Run(":9090"))
}

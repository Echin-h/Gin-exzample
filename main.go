package main

import "LearningGo/cmd/server"

func main() {
	server.Init() // DB,config的Init
	server.Run()  //log,recovery,router的run
}

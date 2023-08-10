package main

import (
	"diplomski.com/db"
	"diplomski.com/server"
	"diplomski.com/utils"
)

func main() {
	utils.LoadEnv()
	db.SetupDB()
	server.RunServer()
}

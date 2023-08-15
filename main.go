package main

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/server"
	"github.com/DzoniDiplomski/Backend_API/utils"
)

func main() {
	utils.LoadEnv()
	db.SetupDB()
	server.RunServer()
}

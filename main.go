package main

import (
	"diplomski.com/db"
	"diplomski.com/server"
)

func main() {
	db.SetupDB()
	server.RunServer()
}

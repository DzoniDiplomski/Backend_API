package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DBConn *sql.DB

func SetupDB() {
	var err error
	DBConn, err = sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Check if the connection is successful
	if err = DBConn.Ping(); err != nil {
		DBConn.Close()
		fmt.Print(err)
		os.Exit(1)
	}
}

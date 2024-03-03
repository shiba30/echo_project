package main

import (
	"fmt"

	"twitter_echo_backend/db"
	"twitter_echo_backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}

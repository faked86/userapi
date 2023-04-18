package main

import (
	"userapi/pkg/db"
	"userapi/pkg/server"
)

func main() {
	port := "8000"
	storageFile := "users.json"

	db := db.NewUserDB(storageFile)

	s := server.NewServer(port, &db)
	s.Start()
}

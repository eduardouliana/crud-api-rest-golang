package main

import (
	"dudu/go-rest/database"
	"dudu/go-rest/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}

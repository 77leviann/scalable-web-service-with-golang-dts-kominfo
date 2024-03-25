package main

import (
	"mygram/database"
	"mygram/router"
)

func main() {
	database.StartDB()
	router.StartServer()
}

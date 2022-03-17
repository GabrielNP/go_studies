package main

import (
	"github.com/joho/godotenv"
	
	"gin_and_tests/database"
	"gin_and_tests/routes"
)

func main() {
	godotenv.Load()
	database.Connect()
	routes.HandleRequests()
}

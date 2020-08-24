package main

import (
	"github.com/Zaysevkun/RESTful-API/api"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// connect .env file and read config variables
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
	config := api.NewConfig()

	// start server
	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}

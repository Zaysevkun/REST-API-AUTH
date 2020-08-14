package main

import (
	"github.com/Zaysevkun/RESTful-API/api"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
	config := api.NewConfig()
	s := api.New(config)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}

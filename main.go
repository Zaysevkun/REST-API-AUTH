package main

import (
	"./api"
	"log"
	"os"
)

func main() {
	config := api.NewConfig()
	s := api.New(config)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}

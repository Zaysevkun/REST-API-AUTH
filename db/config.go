package db

import "os"

type Config struct {
	DatabaseUrl string
}

func NewConfig() *Config {
	dburl := os.Getenv("db_url")

	return &Config{
		DatabaseUrl: dburl,
	}
}

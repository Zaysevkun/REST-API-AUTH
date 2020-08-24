package api

import "os"

// Server config
type Config struct {
	Port        string
	LogLevel    string
	DatabaseURL string
}

// Generate Config from .env
func NewConfig() *Config {
	var c Config
	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		c.Port = ":8000"
	}
	c.LogLevel = os.Getenv("LogLevel")
	if c.LogLevel == "" {
		c.LogLevel = "debug"
	}
	c.DatabaseURL = os.Getenv("DATABASE_URL")
	if c.DatabaseURL == "" {
		c.DatabaseURL = "host=localhost user=postgres dbname=rest_api sslmode=disable password=4256"
	}
	return &c
}

package api

import "os"

type ApiServer struct {
}

func New(config Config) *ApiServer {
	return &ApiServer{}
}

func (s *ApiServer) Start() error {
	return nil
}

type Config struct {
	Port string
}

func NewConfig() Config {
	var c Config
	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		c.Port = "8000"
	}
	return c
}

package api

import (
	"github.com/Zaysevkun/RESTful-API/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

type ApiServer struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *db.Storage
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.cofigureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	if err := s.configureStorage(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.Port, s.router)
}

type Config struct {
	Port     string
	LogLevel string
}

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
	return &c
}

func (s *ApiServer) cofigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *ApiServer) configureStorage() error {
	db := db.New(db.NewConfig())
	if err := db.Open(); err != nil {
		return err
	}
	s.storage = db
	return nil
}

func (s *ApiServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "HELLO")
	}
}

package server

import (
	"github.com/Heroin-lab/diplom.git/internal/app/database"
	"github.com/Heroin-lab/diplom.git/internal/app/server/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	config   *Config
	logger   *logrus.Logger
	router   *mux.Router
	dataBase *database.Database
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureDatabase(); err != nil {
		return err
	}

	s.logger.Info("Starting server!")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/hello", handlers.HandleHello())
}

func (s *Server) configureDatabase() error {
	db := database.New(s.config.DatabaseConfig)
	if err := db.Open(); err != nil {
		return err
	}

	s.dataBase = db

	return nil
}

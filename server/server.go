package server

import (
	"errors"
	"log"
	"mutants-meli/internal/database"
	"mutants-meli/internal/handlers"
	"mutants-meli/internal/models"
	"mutants-meli/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(config *models.Config) error {
	if config.Server_port == "" {
		return errors.New("port is required")
	}
	if config.Db_host == "" {
		return errors.New("db_host is required")
	}

	r := mux.NewRouter()
	BindRoutes(r)
	http.Handle("/", r)

	repo, err := database.NewDatabaseRepository(config)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	server := http.Server{Addr: config.Server_port}
	log.Println("Starting server on port", config.Server_port)
	err = server.ListenAndServe()

	return err
}

func BindRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler()).Methods(http.MethodGet)
	r.HandleFunc("/mutant", handlers.MutantHandler()).Methods(http.MethodPost)
	r.HandleFunc("/stats", handlers.StatsHandler()).Methods(http.MethodGet)
}

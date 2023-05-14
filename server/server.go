package server

import (
	"errors"
	"mutants-meli/internal/database"
	"mutants-meli/internal/handlers"
	"mutants-meli/internal/models"
	"mutants-meli/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(config *models.Config) (http.Server, error) {
	err := ValidateConfig(config)
	if err != nil {
		return http.Server{}, err
	}

	repo, err := database.NewDatabaseRepository(config)
	if err != nil {
		return http.Server{}, err
	}
	repository.SetRepository(repo)

	server := http.Server{Addr: config.Server_port}

	return server, nil
}

func BindRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/mutant", handlers.MutantHandler).Methods(http.MethodPost)
	r.HandleFunc("/stats", handlers.StatsHandler).Methods(http.MethodGet)
	http.Handle("/", r)
}

func ValidateConfig(config *models.Config) error {
	if config.Server_port == "" {
		return errors.New("port is required")
	}
	if config.Db_host == "" {
		return errors.New("db host is required")
	}
	if config.Db_user == "" {
		return errors.New("user is required")
	}
	if config.Db_pass == "" {
		return errors.New("password is required")
	}
	if config.Db_name == "" {
		return errors.New("db name is required")
	}

	return nil
}

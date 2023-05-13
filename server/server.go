package server

import (
	"context"
	"errors"
	"log"
	"mutants-meli/internal/database"
	"mutants-meli/internal/handlers"
	"mutants-meli/internal/models"
	"mutants-meli/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func NewServer(config *models.Config) error {
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

	r := mux.NewRouter()
	BindRoutes(r)
	http.Handle("/", r)

	repo, err := database.NewDatabaseRepository(config)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	server := http.Server{Addr: config.Server_port}

	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = server.ListenAndServe()
	}()
	log.Println("Starting server on port", config.Server_port)
	<-serverDoneChan

	server.Shutdown(context.Background())
	log.Println("Server stopped")

	return err
}

func BindRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler()).Methods(http.MethodGet)
	r.HandleFunc("/mutant", handlers.MutantHandler()).Methods(http.MethodPost)
	r.HandleFunc("/stats", handlers.StatsHandler()).Methods(http.MethodGet)
}

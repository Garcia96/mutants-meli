package main

import (
	"context"
	"log"
	"os"

	"mutants-meli/internal/models"
	"mutants-meli/server"

	"github.com/joho/godotenv"
)

func main() {
	config := SetupServerConfig()

	s, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	server.BindRoutes()

	log.Println("Starting server on port", config.Server_port)
	err = s.ListenAndServe()
	if err != nil {
		s.Shutdown(context.Background())
		log.Fatal(err)
	}
}

func SetupServerConfig() *models.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server_port := os.Getenv("port")
	user := os.Getenv("db_user")
	pass := os.Getenv("db_pass")
	host := os.Getenv("db_host")
	db_name := os.Getenv("db_name")
	port := os.Getenv("db_port")

	return &models.Config{
		Server_port: server_port,
		Db_host:     host,
		Db_name:     db_name,
		Db_port:     port,
		Db_user:     user,
		Db_pass:     pass,
	}
}

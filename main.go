package main

import (
	"log"
	"os"

	"mutants-meli/internal/models"
	"mutants-meli/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server_port := os.Getenv("server_port")
	user := os.Getenv("user")
	pass := os.Getenv("pass")
	host := os.Getenv("host")
	db_name := os.Getenv("db_name")
	port := os.Getenv("port")

	err = server.NewServer(&models.Config{Server_port: server_port, Db_host: host, Db_name: db_name, Db_port: port, Db_user: user, Db_pass: pass})
	if err != nil {
		log.Fatal(err)
	}

	// s.Start(BindRoutes)

}
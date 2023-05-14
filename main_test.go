package main

import (
	"mutants-meli/internal/models"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ServerMock struct {
	mock.Mock
}

func (s *ServerMock) NewServer(c *models.Config) error {
	args := s.Called(c)
	return args.Error(0)
}

func TestSetupServerConfig(t *testing.T) {
	c := require.New(t)
	os.Setenv("port", "8080")
	os.Setenv("db_user", "user")
	os.Setenv("db_pass", "pass")
	os.Setenv("db_host", "localhost")
	os.Setenv("db_name", "mydb")
	os.Setenv("db_port", "5432")

	server_port := os.Getenv("port")
	user := os.Getenv("db_user")
	pass := os.Getenv("db_pass")
	host := os.Getenv("db_host")
	db_name := os.Getenv("db_name")
	port := os.Getenv("db_port")

	config := SetupServerConfig()

	c.Equal(host, config.Db_host)
	c.Equal(server_port, config.Server_port)
	c.Equal(user, config.Db_user)
	c.Equal(pass, config.Db_pass)
	c.Equal(db_name, config.Db_name)
	c.Equal(port, config.Db_port)

	os.Unsetenv("port")
	os.Unsetenv("db_user")
	os.Unsetenv("db_pass")
	os.Unsetenv("db_host")
	os.Unsetenv("db_name")
	os.Unsetenv("db_port")
}

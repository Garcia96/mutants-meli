package server

import (
	"errors"
	"mutants-meli/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestServerSuccess(t *testing.T) {
	c := require.New(t)
	config := &models.Config{
		Server_port: ":8080",
		Db_host:     "localhost",
		Db_name:     "database",
		Db_port:     "3306",
		Db_user:     "user",
		Db_pass:     "pass",
	}

	s, err := NewServer(config)
	c.Equal(nil, err)
	defer s.Close()
}

func TestBindRoutes(t *testing.T) {
	c := require.New(t)
	r := mux.NewRouter()

	BindRoutes()

	req1, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := httptest.NewRecorder()
	r.ServeHTTP(rr1, req1)
	c.NotEqual(http.StatusOK, rr1.Code)
}

func TestValidateConfigServerPort(t *testing.T) {
	c := require.New(t)
	var expectedError = errors.New("port is required")
	config := &models.Config{
		Server_port: "",
		Db_host:     "localhost",
		Db_name:     "database",
		Db_port:     "3306",
		Db_user:     "user",
		Db_pass:     "pass",
	}

	s, err := NewServer(config)
	c.EqualError(expectedError, err.Error())

	defer s.Close()
}

func TestValidateConfigDbHost(t *testing.T) {
	c := require.New(t)
	var expectedError = errors.New("db host is required")
	config := &models.Config{
		Server_port: ":8080",
		Db_host:     "",
		Db_name:     "database",
		Db_port:     "3306",
		Db_user:     "user",
		Db_pass:     "pass",
	}

	err := ValidateConfig(config)

	c.EqualError(expectedError, err.Error())
}

func TestValidateConfigDbName(t *testing.T) {
	c := require.New(t)
	var expectedError = errors.New("db name is required")
	config := &models.Config{
		Server_port: ":8080",
		Db_host:     "localhost",
		Db_name:     "",
		Db_port:     "3306",
		Db_user:     "user",
		Db_pass:     "pass",
	}

	err := ValidateConfig(config)

	c.EqualError(expectedError, err.Error())
}

func TestValidateConfigDbUser(t *testing.T) {
	c := require.New(t)
	var expectedError = errors.New("user is required")
	config := &models.Config{
		Server_port: ":8080",
		Db_host:     "localhost",
		Db_name:     "database",
		Db_port:     "3306",
		Db_user:     "",
		Db_pass:     "pass",
	}

	err := ValidateConfig(config)

	c.EqualError(expectedError, err.Error())
}

func TestValidateConfigDbPass(t *testing.T) {
	c := require.New(t)
	var expectedError = errors.New("password is required")
	config := &models.Config{
		Server_port: ":8080",
		Db_host:     "localhost",
		Db_name:     "database",
		Db_port:     "3306",
		Db_user:     "user",
		Db_pass:     "",
	}

	err := ValidateConfig(config)

	c.EqualError(expectedError, err.Error())
}

func TestValidateConfigSuccess(t *testing.T) {
	c := require.New(t)
	config := &models.Config{
		Server_port: ":8080",
		Db_host:     "localhost",
		Db_name:     "database",
		Db_port:     "3306",
		Db_user:     "user",
		Db_pass:     "pass",
	}

	s, err := NewServer(config)

	c.Equal(nil, err)
	defer s.Close()
}

package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"mutants-meli/internal/models"
	"mutants-meli/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockRepo struct{}

func (m *mockRepo) GetAllDna(ctx context.Context) ([]*models.Dna, error) {
	dnaList := []*models.Dna{
		{Id: "1", Dna_string: `'["ATGCGA","CAGTGC","TTATGT","AGAATG","CCCTTA","TCACTG"]'`, Is_mutant: false},
	}
	return dnaList, nil
}

func (m *mockRepo) InsertDna(ctx context.Context, dna *models.Dna) error {
	return nil
}

type mockRepoError struct{}

func (m *mockRepoError) GetAllDna(ctx context.Context) ([]*models.Dna, error) {
	dnaList := []*models.Dna{}
	return dnaList, errors.New("cannot get dnas")
}

func (m *mockRepoError) InsertDna(ctx context.Context, dna *models.Dna) error {
	return errors.New("cannot insert dnas")
}

func TestHomeHandler(t *testing.T) {
	var expectedReponse = models.Response{
		Message: "Welcome Mutants-meli!",
		Status:  true,
	}
	var response models.Response
	c := require.New(t)

	request, err := http.NewRequest("GET", "/", nil)
	c.NoError(err)
	w := httptest.NewRecorder()

	HomeHandler(w, request)

	err = json.Unmarshal([]byte(w.Body.Bytes()), &response)
	c.NoError(err)

	c.Equal(http.StatusOK, w.Code)
	c.Equal(expectedReponse, response)
}

func TestStatsHandlerOk(t *testing.T) {
	var response models.StatsResponse
	var expectedReponse = models.StatsResponse{
		Count_mutant_dna: 0,
		Count_human_dna:  1,
		Ratio:            0,
	}

	c := require.New(t)

	request, err := http.NewRequest("GET", "/stats", nil)
	c.NoError(err)
	w := httptest.NewRecorder()

	repo := &mockRepo{}
	repository.SetRepository(repo)

	StatsHandler(w, request)

	err = json.Unmarshal([]byte(w.Body.Bytes()), &response)
	c.NoError(err)

	c.Equal(http.StatusOK, w.Code)
	c.Equal(expectedReponse, response)
}

func TestStatsHandlerError(t *testing.T) {
	c := require.New(t)

	request, err := http.NewRequest("GET", "/stats", nil)
	c.NoError(err)
	w := httptest.NewRecorder()

	repo := &mockRepoError{}
	repository.SetRepository(repo)

	StatsHandler(w, request)

	c.Equal(http.StatusInternalServerError, w.Code)
}

func TestMutantHandlerOK(t *testing.T) {
	c := require.New(t)

	body := []byte(`{"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)

	request, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(body))
	c.NoError(err)
	w := httptest.NewRecorder()

	repo := &mockRepo{}
	repository.SetRepository(repo)

	MutantHandler(w, request)

	c.Equal(http.StatusOK, w.Code)
}

func TestMutantHandlerBadRequest(t *testing.T) {
	c := require.New(t)

	body := []byte(`{"dna": "TEST"`)

	request, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(body))
	c.NoError(err)
	w := httptest.NewRecorder()

	repo := &mockRepo{}
	repository.SetRepository(repo)

	MutantHandler(w, request)

	c.Equal(http.StatusBadRequest, w.Code)
}

func TestMutantHandlerBadDnaString(t *testing.T) {
	c := require.New(t)

	body := []byte(`{"dna": ["ATGCGA","CAGTXC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)

	request, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(body))
	c.NoError(err)
	w := httptest.NewRecorder()

	repo := &mockRepo{}
	repository.SetRepository(repo)

	MutantHandler(w, request)

	var response string
	err = json.Unmarshal([]byte(w.Body.Bytes()), &response)
	c.NoError(err)

	c.Equal(http.StatusBadRequest, w.Code)
	c.Equal("dna string is invalid", response)
}

func TestMutantHandlerError(t *testing.T) {
	c := require.New(t)

	body := []byte(`{"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)

	request, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(body))
	c.NoError(err)
	w := httptest.NewRecorder()

	repo := &mockRepoError{}
	repository.SetRepository(repo)

	MutantHandler(w, request)

	c.Equal(http.StatusInternalServerError, w.Code)
}

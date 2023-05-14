package repository

import (
	"context"
	"mutants-meli/internal/models"
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

func TestGetAllDna(t *testing.T) {
	c := require.New(t)
	expectedReponse := []*models.Dna{
		{
			Id:         "1",
			Dna_string: `'["ATGCGA","CAGTGC","TTATGT","AGAATG","CCCTTA","TCACTG"]'`,
			Is_mutant:  false,
		},
	}

	repo := &mockRepo{}
	SetRepository(repo)

	dnas, err := GetAllDna(context.Background())
	c.NoError(err)

	c.Equal(expectedReponse, dnas)
}

func TestInsertDna(t *testing.T) {
	c := require.New(t)
	repo := &mockRepo{}
	SetRepository(repo)

	dna := models.Dna{
		Id:         "1",
		Dna_string: `'["ATGCGA","CAGTGC","TTATGT","AGAATG","CCCTTA","TCACTG"]'`,
		Is_mutant:  false,
	}

	err := InsertDna(context.Background(), &dna)
	c.NoError(err)
}

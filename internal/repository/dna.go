package repository

import (
	"context"
	"mutants-meli/internal/models"
)

type DnaRepository interface {
	GetAllDna(ctx context.Context) ([]*models.Dna, error)
	InsertDna(ctx context.Context, dna *models.Dna) error
}

var implementation DnaRepository

func SetRepository(repository DnaRepository) {
	implementation = repository
}

func InsertDna(ctx context.Context, dna *models.Dna) error {
	return implementation.InsertDna(ctx, dna)
}

func GetAllDna(ctx context.Context) ([]*models.Dna, error) {
	return implementation.GetAllDna(ctx)
}

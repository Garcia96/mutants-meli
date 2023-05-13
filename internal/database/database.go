package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mutants-meli/internal/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseRepository struct {
	Db *sql.DB
}

func NewDatabaseRepository(config *models.Config) (*DatabaseRepository, error) {
	var url = fmt.Sprintf("%s:%s@/%s", config.Db_user, config.Db_pass, config.Db_name)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	return &DatabaseRepository{db}, nil
}

func (repo *DatabaseRepository) InsertDna(ctx context.Context, dna *models.Dna) error {
	_, err := repo.Db.ExecContext(ctx, "INSERT INTO dna (id,dna_string,is_mutant, created_at) VALUES (?, ?, ?, ?)", dna.Id, dna.Dna_string, dna.Is_mutant, time.Now())
	return err
}

func (repo *DatabaseRepository) GetAllDna(ctx context.Context) ([]*models.Dna, error) {
	rows, err := repo.Db.QueryContext(ctx, "SELECT id, dna_string, is_mutant FROM dna ")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	dnas := []*models.Dna{}

	for rows.Next() {
		dna := models.Dna{}
		if err = rows.Scan(&dna.Id, &dna.Dna_string, &dna.Is_mutant); err == nil {
			dnas = append(dnas, &dna)
		}
	}

	return dnas, err
}

func (repo *DatabaseRepository) Close() error {
	return repo.Db.Close()
}

package database

import (
	"context"
	"database/sql"
	"mutants-meli/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

var config = models.Config{
	Server_port: ":8080",
	Db_host:     "localhost",
	Db_name:     "test",
	Db_port:     "3306",
	Db_user:     "admin",
	Db_pass:     "test",
}

var dna = models.Dna{
	Id:         "1",
	Dna_string: `'["ATGCGA","CAGTGC","TTATGT","AGAATG","CCCTTA","TCACTG"]'`,
	Is_mutant:  false,
}

func TestNewDatabaseRepository(t *testing.T) {
	c := require.New(t)
	db, mock, err := sqlmock.New()
	c.NoError(err)
	defer db.Close()

	mock.ExpectPing()

	repo, err := NewDatabaseRepository(&config)
	c.NoError(err)

	c.IsType(&sql.DB{}, repo.Db)
}

func TestInsertDna(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.NoError(err)
	defer db.Close()

	repo := &DatabaseRepository{db}

	mock.ExpectExec("INSERT INTO dna").WithArgs(dna.Id, dna.Dna_string, dna.Is_mutant, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertDna(context.Background(), &dna)
	c.NoError(err)
}

func TestGetAllDna(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.NoError(err)
	defer db.Close()

	repo := &DatabaseRepository{db}

	rows := mock.NewRows([]string{"id", "dna_string", "is_mutant"}).AddRow(dna.Id, dna.Dna_string, dna.Is_mutant)
	mock.ExpectQuery("SELECT id, dna_string, is_mutant FROM dna").WillReturnRows(rows)

	dnas, err := repo.GetAllDna(context.Background())
	c.NoError(err)

	expectedResponse := []*models.Dna{
		{
			Id:         "1",
			Dna_string: `'["ATGCGA","CAGTGC","TTATGT","AGAATG","CCCTTA","TCACTG"]'`,
			Is_mutant:  false,
		},
	}

	c.Equal(expectedResponse, dnas)
}

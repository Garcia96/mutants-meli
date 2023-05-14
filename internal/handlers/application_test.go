package handlers

import (
	"encoding/json"
	"mutants-meli/internal/models"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var request models.MutantRequest
var dnaValid = []byte(`{"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)
var dnaInvalid = []byte(`{"dna": ["ATXCGA","CAGTGC","TTATGT","AGAATG","CCCTTA","TCACTG"]}`)
var expectedDnaModel = models.Dna{
	Id:         "",
	Dna_string: `["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]`,
	Is_mutant:  true,
}

func TestValidateDnaSuccess(t *testing.T) {
	c := require.New(t)

	err := json.Unmarshal(dnaValid, &request)
	c.NoError(err)

	res := ValidateDna(request.Dna)

	c.Equal(true, res)
}

func TestValidateDnaError(t *testing.T) {
	c := require.New(t)

	err := json.Unmarshal(dnaInvalid, &request)
	c.NoError(err)

	res := ValidateDna(request.Dna)

	c.Equal(false, res)
}

func TestIsMutantSuccessOk(t *testing.T) {
	c := require.New(t)

	err := json.Unmarshal(dnaValid, &request)
	c.NoError(err)

	dnaModel, status := IsMutant(request.Dna)

	c.Equal(expectedDnaModel.Dna_string, dnaModel.Dna_string)
	c.Equal(http.StatusOK, status)
}

func TestIsMutantSuccessForbidden(t *testing.T) {
	c := require.New(t)

	err := json.Unmarshal(dnaInvalid, &request)
	c.NoError(err)

	dnaModel, status := IsMutant(request.Dna)

	c.NotEqual(expectedDnaModel.Dna_string, dnaModel.Dna_string)
	c.Equal(http.StatusForbidden, status)
}

func TestVerifyStatsSuccess(t *testing.T) {
	c := require.New(t)

	body, err := os.ReadFile("samples/dna_response.json")
	c.NoError(err)

	var request = []*models.Dna{}
	err = json.Unmarshal([]byte(body), &request)
	c.NoError(err)

	var expectedRes = models.StatsResponse{
		Count_mutant_dna: 3,
		Count_human_dna:  4,
		Ratio:            0.75,
	}

	res := VerifyStats(request)

	c.Equal(expectedRes, res)
}

func TestVerifyStatsError(t *testing.T) {
	c := require.New(t)
	var expectedRes = models.StatsResponse{
		Count_mutant_dna: 1,
		Count_human_dna:  1,
		Ratio:            1,
	}

	var dnas = []*models.Dna{}
	var dna = models.Dna{
		Id:         "",
		Dna_string: `["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]`,
		Is_mutant:  true,
	}
	dnas = append(dnas, &dna)

	res := VerifyStats(dnas)
	c.Equal(expectedRes, res)
}

func TestVerifyStatsEmptyDnas(t *testing.T) {
	c := require.New(t)
	var expectedRes = models.StatsResponse{
		Count_mutant_dna: 0,
		Count_human_dna:  0,
		Ratio:            0,
	}

	var dnas = []*models.Dna{}

	res := VerifyStats(dnas)
	c.Equal(expectedRes, res)
}

package handlers

import (
	"encoding/json"
	"log"
	"mutants-meli/internal/models"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func IsMutant(dna []string) (*models.Dna, int) {
	bytes, err := json.Marshal(dna)
	if err != nil {
		log.Fatal(err.Error())
	}
	str := string(bytes)

	is_mutant := CheckMutant(dna)

	dnaModel := models.Dna{
		Id:         uuid.New().String(),
		Dna_string: str,
		Is_mutant:  is_mutant,
	}

	status := func() int {
		if is_mutant {
			return http.StatusOK
		} else {
			return http.StatusForbidden
		}
	}()

	return &dnaModel, status
}

func CheckMutant(dna []string) bool {
	validC := []string{"A", "T", "C", "G"}
	count := 0
	matrix := ConvertToMatrix(dna)

	for i := 0; i < len(validC); i++ {
		if CheckMatrix(matrix, validC[i]) {
			count++
		}
	}

	return count > 1
}

func CheckMatrix(matrix [][]string, sequence string) bool {
	for i := 0; i < len(matrix); i++ {
		count := 0
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == sequence {
				count++
				if count == 4 {
					return true
				}
			} else {
				count = 0
			}
		}
	}

	// Check vertical sequences
	for i := 0; i < len(matrix[0]); i++ {
		count := 0
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == sequence {
				count++
				if count == 4 {
					return true
				}
			} else {
				count = 0
			}
		}
	}

	// Check diagonal sequences
	for i := 0; i < len(matrix)-3; i++ {
		for j := 0; j < len(matrix[0])-3; j++ {
			if matrix[i][j] == sequence && matrix[i+1][j+1] == sequence && matrix[i+2][j+2] == sequence && matrix[i+3][j+3] == sequence {
				return true
			}
			if matrix[i][j+3] == sequence && matrix[i+1][j+2] == sequence && matrix[i+2][j+1] == sequence && matrix[i+3][j] == sequence {
				return true
			}
		}
	}

	return false
}

func ConvertToMatrix(dna []string) [][]string {
	w := len(dna[0])
	h := len(dna)

	matrix := make([][]string, h)
	for i := range matrix {
		matrix[i] = make([]string, w)
	}

	for i, row := range dna {
		for j, cl := range row {
			matrix[i][j] = string(cl)
		}
	}

	return matrix
}

func ValidateDna(dna []string) bool {
	allowedChars := "ATCG"

	for _, str := range dna {
		for _, char := range str {
			if !strings.ContainsRune(allowedChars, char) {
				return false
			}
		}
	}

	return true
}

func VerifyStats(dnas []*models.Dna) models.StatsResponse {
	var count_mutant_dna, count_human_dna int
	var stats models.StatsResponse
	for _, dna := range dnas {
		if dna.Is_mutant {
			count_mutant_dna++
		}
	}

	count_human_dna = len(dnas)
	if count_human_dna == 0 {
		stats = models.StatsResponse{
			Count_mutant_dna: count_mutant_dna,
			Count_human_dna:  count_human_dna,
			Ratio:            0,
		}
		return stats
	}

	ratio := float64(count_mutant_dna) / float64(count_human_dna)

	log.Println("RATIO en Verify", ratio)
	stats = models.StatsResponse{
		Count_mutant_dna: count_mutant_dna,
		Count_human_dna:  count_human_dna,
		Ratio:            ratio,
	}

	return stats
}

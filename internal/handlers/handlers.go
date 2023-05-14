package handlers

import (
	"encoding/json"
	"mutants-meli/internal/models"
	"mutants-meli/internal/repository"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "Welcome Mutants-meli!",
		Status:  true,
	})
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dnas, err := repository.GetAllDna(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	response := VerifyStats(dnas)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func MutantHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request = models.MutantRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isValidDna := ValidateDna(request.Dna)
	if !isValidDna {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("dna string is invalid")
		return
	}

	dnaModel, status := IsMutant(request.Dna)

	err = repository.InsertDna(r.Context(), dnaModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(status)
}

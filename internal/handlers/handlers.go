package handlers

import (
	"encoding/json"
	"mutants-meli/internal/models"
	"mutants-meli/internal/repository"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Message: "Welcome Mutants-meli!",
			Status:  true,
		})
	}
}

func StatsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		dnas, err := repository.GetAllDna(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		response, err := VerifyStats(dnas)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func MutantHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request = models.MutantRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		isInvalidDna := ValidateDna(request.Dna)
		if isInvalidDna {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dnaModel, status := IsMutant(request.Dna)
		w.WriteHeader(status)

		err = repository.InsertDna(r.Context(), dnaModel)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
	}
}

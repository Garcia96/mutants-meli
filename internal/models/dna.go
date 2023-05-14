package models

type Dna struct {
	Id         string `json:"id"`
	Dna_string string `json:"dna_string"`
	Is_mutant  bool   `json:"is_mutant"`
}

type MutantRequest struct {
	Dna []string `json:"dna"`
}

type StatsResponse struct {
	Count_mutant_dna int     `json:"count_mutant_dna"`
	Count_human_dna  int     `json:"count_human_dna"`
	Ratio            float64 `json:"ratio"`
}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

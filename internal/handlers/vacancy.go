package handlers

import (
	"eSearcher/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

func (r *Router) CreateVacancy(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var vacancy *models.Vacancy
	err = json.Unmarshal(body, &vacancy)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}
	if err = r.Services.VacancyService.CreateVacancy(vacancy.Name, vacancy.Description); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

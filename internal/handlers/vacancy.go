package handlers

import (
	"eSearcher/internal/models"
	"encoding/json"
	"fmt"
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
	fmt.Println(vacancy.Tags)
	if err = r.Services.VacancyService.CreateVacancy(vacancy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (r *Router) KeyWordSearchVacancy(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var params *models.SearchVacancyParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}
	fmt.Printf("params to search vacancy: %+v\n", params)
	vacancies, err := r.Services.VacancyService.SearchVacancy(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(vacancies) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := json.Marshal(vacancies)
	if err != nil {
		http.Error(w, "cannot marshall data", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

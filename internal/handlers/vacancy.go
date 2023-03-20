package handlers

import (
	"eSearcher/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (r *Router) GetMyVacancies(w http.ResponseWriter, req *http.Request) {
	userID, err := r.getUserIDFromToken(w, req)
	if err != nil {
		return
	}
	vacancies, err := r.Services.VacancyService.GetEmployerVacancies(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(vacancies)
	if err != nil {
		http.Error(w, "cannot marshall data", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (r *Router) CreateVacancy(w http.ResponseWriter, req *http.Request) {
	//userID, err := r.getUserIDFromToken(w, req)
	//if err != nil {
	//	return
	//}
	//userRole, err := r.getUserRoleFromToken(w, req)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("userID: %v, userRole: %v", userID, userRole)
	//if userRole != 3 {
	//	http.Error(w, "you are not employer", http.StatusInternalServerError)
	//	return
	//}

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

func (r *Router) UpdateVacancy(w http.ResponseWriter, req *http.Request) {
	//userID, err := r.getUserIDFromToken(w, req)
	//if err != nil {
	//	return
	//}
	//userRole, err := r.getUserRoleFromToken(w, req)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("userID: %v, userRole: %v", userID, userRole)
	//if userRole != 3 {
	//	http.Error(w, "you are not employer", http.StatusInternalServerError)
	//	return
	//}

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
	fmt.Printf("%+v\n", vacancy)
	if err = r.Services.VacancyService.UpdateVacancy(vacancy); err != nil {
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

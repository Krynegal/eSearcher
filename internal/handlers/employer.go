package handlers

import (
	"eSearcher/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func (r *Router) GetEmployer(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	employer, err := r.Services.EmployersService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v", employer)
	res, err := json.Marshal(employer)
	if err != nil {
		http.Error(w, "cannot marshall data", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (r *Router) CreateEmployer(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var employer *models.Employer
	err = json.Unmarshal(body, &employer)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v", employer)
	if err = r.Services.EmployersService.Create(employer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (r *Router) UpdateEmployer(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var employer *models.Employer
	err = json.Unmarshal(body, &employer)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v", employer)
	if err = r.Services.EmployersService.Update(employer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

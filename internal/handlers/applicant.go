package handlers

import (
	"eSearcher/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (r *Router) CreateApplicant(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var applicant *models.Applicant
	err = json.Unmarshal(body, &applicant)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}
	fmt.Println(applicant)
	if err = r.Services.ApplicantsService.Create(applicant); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (r *Router) SearchApplicant(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var params *models.SearchApplicantParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}
	fmt.Printf("params to search applicant: %+v\n", params)
	applicants, err := r.Services.ApplicantsService.SearchApplicant(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(applicants) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := json.Marshal(applicants)
	if err != nil {
		http.Error(w, "cannot marshall data", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

package handlers

import (
	"eSearcher/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

func (r *Router) AddResponse(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var response *models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}

	if err = r.Services.ResponsesService.Add(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (r *Router) DeleteResponse(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "incorrect input data", http.StatusInternalServerError)
		return
	}
	var response *models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		http.Error(w, "cannot unmarshall data", http.StatusInternalServerError)
		return
	}

	if err = r.Services.ResponsesService.Delete(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

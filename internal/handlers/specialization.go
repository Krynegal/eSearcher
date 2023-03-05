package handlers

import (
	"net/http"
)

func (r *Router) Get(w http.ResponseWriter, req *http.Request) {
	if _, err := r.Services.SpecializationsService.Get(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

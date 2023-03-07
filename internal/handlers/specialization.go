package handlers

import (
	"encoding/json"
	"net/http"
)

func (r *Router) GetAllSpecializations(w http.ResponseWriter, req *http.Request) {
	specializations, err := r.Services.SpecializationsService.GetAllSpecializations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(specializations)
	if err != nil {
		http.Error(w, "cannot marshall data", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

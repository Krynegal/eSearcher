package handlers

import (
	"encoding/json"
	"net/http"
)

func (r *Router) GetAllOptions(w http.ResponseWriter, req *http.Request) {
	v := req.URL.Query()
	names := v["name"]

	options, err := r.Services.OptionsService.GetAll(names)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(options)
	if err != nil {
		http.Error(w, "cannot marshall data", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/repositories"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	prj, err := repositories.GetAllProjects()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(prj)
	}
}

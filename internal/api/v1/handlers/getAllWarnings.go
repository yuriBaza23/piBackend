package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/repositories"
)

func GetAllWarnings(w http.ResponseWriter, r *http.Request) {
	war, err := repositories.GetAllWarnings()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(war)
	}
}

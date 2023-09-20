package handlers

import (
	"encoding/json"
	"net/http"
	"pi/cmd/internal/api/v1/models"
	"pi/cmd/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func UpdateWarning(w http.ResponseWriter, r *http.Request) {
	var input models.Warning

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	war, err := repositories.GetWarning(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Title == "" {
		input.Title = war.Title
	}
	if input.Content == "" {
		input.Content = war.Content
	}

	_, err = repositories.UpdateWarning(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	war, err = repositories.GetWarning(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(war)
}

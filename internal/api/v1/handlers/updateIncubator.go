package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func UpdateIncubator(w http.ResponseWriter, r *http.Request) {
	var input models.Incubator

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	inc, err := repositories.GetIncubator(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		input.Name = inc.Name
	}
	if input.Email == "" {
		input.Email = inc.Email
	}
	if input.Password != "" {
		hashPassword, err := repositories.HashPassword(input.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		input.Password = hashPassword
	} else {
		input.Password = inc.Password
	}

	_, err = repositories.UpdateIncubator(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inc, err = repositories.GetIncubator(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inc)
}

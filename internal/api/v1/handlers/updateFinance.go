package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateFinance(w http.ResponseWriter, r *http.Request) {
	var input models.Finance

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fin, err := repositories.GetFinance(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		input.Name = fin.Name
	}
	if input.Type == "" {
		input.Type = fin.Type
	}
	if input.FinValue == "" {
		input.Value = fin.Value
	} else {
		input.Value, _ = strconv.Atoi(input.FinValue)
	}

	_, err = repositories.UpdateFinance(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fin, err = repositories.GetFinance(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fin)
}

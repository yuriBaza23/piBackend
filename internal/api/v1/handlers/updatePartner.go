package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func UpdatePartner(w http.ResponseWriter, r *http.Request) {
	var input models.Partner

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	partner, err := repositories.GetPartner(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		input.Name = partner.Name
	}
	if input.Type == "" {
		input.Type = partner.Type
	}
	if input.Email == "" {
		input.Email = partner.Email
	}
	if input.AccountID == "" {
		input.AccountID = partner.AccountID
	}

	_, err = repositories.UpdatePartner(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	partner, err = repositories.GetPartner(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partner)
}

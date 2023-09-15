package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var input models.Company

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cmp, err := repositories.GetCompany(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		input.Name = cmp.Name
	}
	if input.Email == "" {
		input.Email = cmp.Email
	}
	if input.CNPJ == "" {
		input.CNPJ = cmp.CNPJ
	}
	if input.HubID == "" {
		input.HubID = cmp.HubID
	}

	_, err = repositories.UpdateCompany(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmp, err = repositories.GetCompany(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cmp)
}

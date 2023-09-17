package handlers

import (
	"encoding/json"
	"net/http"
	"pi/cmd/internal/api/v1/models"
	"pi/cmd/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var input models.Project

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	prj, err := repositories.GetProject(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		input.Name = prj.Name
	}
	if input.CmpID == "" {
		input.CmpID = prj.CmpID
	}
	if input.Description == "" {
		input.Description = prj.Description
	}

	_, err = repositories.UpdateProject(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prj, err = repositories.GetProject(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prj)
}

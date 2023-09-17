package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input models.Tasks

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tsk, err := repositories.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Title == "" {
		input.Title = tsk.Title
	}
	if input.Description == "" {
		input.Description = tsk.Description
	}
	if input.Status == "" {
		input.Status = tsk.Status
	}
	if input.InitialDate == "" {
		input.InitialDate = tsk.InitialDate
	}
	if input.EndDate == "" {
		input.EndDate = tsk.EndDate
	}

	_, err = repositories.UpdateTask(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tsk, err = repositories.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tsk)
}

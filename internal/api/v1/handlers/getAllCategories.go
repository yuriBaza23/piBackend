package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	cats, err := repositories.GetAllCategories(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cats)
	}
}

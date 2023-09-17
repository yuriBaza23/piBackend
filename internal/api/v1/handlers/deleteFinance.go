package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/repositories"

	"github.com/gorilla/mux"
)

func DeleteFinance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fin, err := repositories.DeleteFinance(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fin)
	}
}

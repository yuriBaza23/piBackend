package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var usr models.User
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verifica se o usuário existe no banco de dados User via email e senha
	res, err := repositories.GetUserByEmailAndPassword(usr.Email, usr.Password)
	if err != nil {
		res, err := repositories.GetIncubatorByEmailAndPassword(usr.Email, usr.Password)
		if err == nil {
			resp = map[string]any{
				"error": false,
				"user":  res,
				"type":  "incubator",
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	} else {
		resp = map[string]any{
			"error": false,
			"user":  res,
			"type":  "company",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

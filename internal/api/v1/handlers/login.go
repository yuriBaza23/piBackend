package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var usr models.User
	var inc models.Incubator
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// A ordem de criação do Hash e sua passagem como parâmetro na função GetUserByEmailAndPassword
	// serve para verificar se um token de uma senha passada por um usuário existe ou não no BD
	hashPassword, err := repositories.HashPassword(usr.Password)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
		return
	}

	// Verifica se o usuário existe no banco de dados User via email e senha
	_, err = repositories.GetUserByEmailAndPassword(usr.Email, hashPassword)
	if err != nil {
		// Caso não exista, realiza a busca no banco de dados Incubator

		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		_, err = repositories.GetIncubatorByEmailAndPassword(inc.Email, hashPassword)
		if err == nil {
			resp = map[string]any{
				"error":   false,
				"message": "The User exists and can be logged in",
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "The User exists and can be logged in",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

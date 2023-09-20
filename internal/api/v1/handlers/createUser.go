package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
	"pi/internal/api/v1/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verifica se o usuário já existe no banco de dados
	_, err = repositories.GetUserByEmail(usr.Email)
	if err == nil {
		resp = map[string]any{
			"error":   true,
			"message": "User already exists",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Cria um ID para o usuário (Formato UUIDv4)
	usr.ID = utils.CreateUuid()

	// Cria um hash para a senha do usuário
	// O hash é criado com o repositório HashPassword e pode devolver
	// um erro caso a senha não seja válida. Esse erro é tratado
	// e retornado para o usuário.
	hashPassword, err := repositories.HashPassword(usr.Password)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
		return
	}

	// A senha do usuário é substituída pelo hash
	usr.Password = hashPassword

	// Diz se o usuário é um administrador ou não
	usr.VerifyUserType()

	// Insere o usuário no banco de dados.
	// O repositório InsertUser pode devolver um erro caso o usuário
	// não seja inserido no banco de dados. Esse erro é tratado e
	// retornado para o usuário. Ainda, faz o relacionamento entre
	// o usuário e a empresa.
	id, err := repositories.InsertUser(usr)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": fmt.Sprintf(`Partner created with ID: %s`, id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

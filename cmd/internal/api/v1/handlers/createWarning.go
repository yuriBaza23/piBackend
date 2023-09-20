package handlers

import (
	"encoding/json"
	"net/http"
	"pi/cmd/internal/api/v1/models"
	"pi/cmd/internal/api/v1/repositories"
	"pi/cmd/internal/api/v1/utils"
)

func CreateWarning(w http.ResponseWriter, r *http.Request) {
	var war models.Warning
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&war)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verifica se a advertência já existe no banco de dados
	_, err = repositories.GetWarning(war.ID)
	if err == nil {
		resp = map[string]any{
			"error":   true,
			"message": "Warning already exists",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Cria um ID para a Advertência (Formato UUIDv4)
	war.ID = utils.CreateUuid()

	// Insere a Advertência no banco de dados.
	// O repositório InsertWaring pode devolver um erro caso a advertência
	// não seja inserida no banco de dados. Esse erro é tratado e
	// retornado para o usuário.
	warningID, err := repositories.InsertWarning(war)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "Warning created with ID: " + warningID,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

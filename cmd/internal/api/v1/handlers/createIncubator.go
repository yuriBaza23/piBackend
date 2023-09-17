package handlers

import (
	"encoding/json"
	"net/http"
	"pi/cmd/internal/api/v1/models"
	"pi/cmd/internal/api/v1/repositories"
	"pi/cmd/internal/api/v1/utils"
)

func CreateIncubator(w http.ResponseWriter, r *http.Request) {
	var inc models.Incubator
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&inc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verifica se a incubadora já existe no banco de dados
	_, err = repositories.GetIncubator(inc.ID)
	if err == nil {
		resp = map[string]any{
			"error":   true,
			"message": "Incubator already exists",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Cria um ID para a incubadora (Formato UUIDv4)
	inc.ID = utils.CreateUuid()

	// Insere a incubadora no banco de dados.
	// O repositório InsertIncubator pode devolver um erro caso a incubadora
	// não seja inserida no banco de dados. Esse erro é tratado e
	// retornado para o usuário.
	incubatorID, err := repositories.InsertIncubator(inc)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "Incubator created with ID: " + incubatorID,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

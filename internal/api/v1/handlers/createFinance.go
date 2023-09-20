package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
	"pi/internal/api/v1/utils"
	"strconv"
)

func CreateFinance(w http.ResponseWriter, r *http.Request) {
	var fin models.Finance
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&fin)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Cria um ID para a finança (Formato UUIDv4)
	fin.ID = utils.CreateUuid()

	fin.Value, _ = strconv.Atoi(fin.FinValue)

	// Insere a finança no banco de dados.
	// O repositório InsertIncubator pode devolver um erro caso a finança
	// não seja inserida no banco de dados. Esse erro é tratado e
	// retornado para o usuário.
	finID, err := repositories.InsertFinance(fin)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "Finance created with ID: " + finID,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

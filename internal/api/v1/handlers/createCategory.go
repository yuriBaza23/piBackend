package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
	"pi/internal/api/v1/utils"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Cria um ID para a categoria (Formato UUIDv4)
	cat.ID = utils.CreateUuid()

	// Insere a categoria criada no banco de dados.
	catID, err := repositories.InsertCategory(cat)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "Category created with ID: " + catID,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

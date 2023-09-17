package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
	"pi/internal/api/v1/utils"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var tsk models.Tasks
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&tsk)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Cria um ID para a tarefa (Formato UUIDv4)
	tsk.ID = utils.CreateUuid()

	// Insere a task no banco de dados.
	// O repositório InsertIncubator pode devolver um erro caso a task
	// não seja inserida no banco de dados. Esse erro é tratado e
	// retornado para o usuário.
	tskID, err := repositories.InsertTask(tsk)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "Task created with ID: " + tskID,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

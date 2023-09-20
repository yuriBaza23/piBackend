package handlers

import (
	"encoding/json"
	"net/http"
	"pi/cmd/internal/api/v1/models"
	"pi/cmd/internal/api/v1/repositories"
	"pi/cmd/internal/api/v1/utils"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var prj models.Project
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&prj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verifica se um projeto com o mesmo nome já existe no banco de dados
	_, err = repositories.GetProject(prj.Name)
	if err == nil {
		resp = map[string]any{
			"error":   true,
			"message": "Project already exists",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Cria um ID para a empresa (Formato UUIDv4)
	prj.ID = utils.CreateUuid()

	// Insere o projeto no banco de dados
	// O repositório InsertProject pode devolver um erro caso o projeto
	// não seja inserido no banco de dados. Esse erro é tratado e
	// retornado para o usuário.
	projectId, err := repositories.InsertProject(prj)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "Project created with ID: " + projectId,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

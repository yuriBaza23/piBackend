package handlers

import (
	"encoding/json"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
	"pi/internal/api/v1/utils"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var cmp models.Company
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&cmp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verifica se o owner da empresa já existe no banco de dados
	_, err = repositories.GetUserByEmail(cmp.OwnerEmail)
	if err == nil {
		resp = map[string]any{
			"error":   true,
			"message": "User already exists",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Verifica se a empresa já existe no banco de dados
	_, err = repositories.GetCompanyByCNPJ(cmp.CNPJ)
	if err == nil {
		resp = map[string]any{
			"error":   true,
			"message": "Company already exists",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Cria um ID para a empresa (Formato UUIDv4)
	cmp.ID = utils.CreateUuid()

	// Insere a empresa no banco de dados.
	// O repositório InsertCompany pode devolver um erro caso a empresa
	// não seja inserida no banco de dados. Esse erro é tratado e
	// retornado para o usuário.
	companyId, err := repositories.InsertCompany(cmp)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	}

	// Criação do owner da empresa de acordo com os dados da empresa
	var usr models.User

	// Cria um ID para o usuário (Formato UUIDv4)
	usr.ID = utils.CreateUuid()
	usr.Name = cmp.OwnerName
	usr.Email = cmp.OwnerEmail

	if cmp.IsPreCad {
		usr.Password = "1234567890"
		usr.IsPreReg = true
	}

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
	usr.Type = "owner"

	// Insere o usuário no banco de dados.
	// O repositório InsertUser pode devolver um erro caso o usuário
	// não seja inserido no banco de dados. Esse erro é tratado e
	// retornado para o usuário. Ainda, faz o relacionamento entre
	// o usuário e a empresa.
	usr.CompanyID = companyId
	usr.ID, err = repositories.InsertUser(usr)
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": err.Error(),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": "Company created with ID: " + companyId + " and owner created with ID: " + usr.ID,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

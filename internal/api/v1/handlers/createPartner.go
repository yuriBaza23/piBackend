package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pi/internal/api/v1/models"
	"pi/internal/api/v1/repositories"
	"pi/internal/api/v1/utils"
)

func CreateParner(w http.ResponseWriter, r *http.Request) {
	var partner models.Partner
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Cria um ID para o sócio (Formato UUIDv4)
	partner.ID = utils.CreateUuid()

	// TODO: Validar se a empresa existe pois não há sentido em
	// criar um sócio de uma empresa que não existe
	// companyId, err := repositories.GetCompany(partner.CompanyId)

	// Se o tipo do sócio for admin, o accountId é o mesmo que o companyId
	// Se não, o accountId ficará em branco e será preenchido quando o sócio
	// terminar o cadastro iniciado pelo sócio admin
	if partner.Type == "admin" {
		partner.AccountID = partner.CompanyID
	}

	validatePartnerEmail := utils.ValidateEmail(partner.Email)
	if !validatePartnerEmail {
		resp = map[string]any{
			"error":   true,
			"message": "email is invalid",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	id, err := repositories.InsertPartner(partner)

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

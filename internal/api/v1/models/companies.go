package models

import (
	"errors"
	"pi/internal/api/v1/utils"
)

type Company struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CNPJ      string `json:"cnpj"`
	HubID     string `json:"hubId"`
	IsPreCad  bool   `json:"isPreCad"`
	Users     []User `json:"users"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	// Para criação do owner
	OwnerName  string `json:"ownerName"`
	OwnerEmail string `json:"ownerEmail"`
}

func (cmp *Company) VerifyCompanyEmail() error {
	if !utils.ValidateEmail(cmp.Email) {
		return errors.New("invalid email")
	}
	return nil
}

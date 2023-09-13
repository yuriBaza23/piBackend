package models

import (
	"errors"
	"pi/cmd/internal/api/v1/utils"
)

type Incubator struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsPreCad bool   `json:"isPreCad"`
	// Para utilizar quando tivermos o back unido
	//Companies    []Companies `json:"companies"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (cmp *Incubator) VerifyIncubatorEmail() error {
	if !utils.ValidateEmail(cmp.Email) {
		return errors.New("invalid email")
	}
	return nil
}

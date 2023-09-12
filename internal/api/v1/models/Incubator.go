package models

import (
	"errors"
	"pi/internal/api/v1/utils"
)

type Incubator struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (inc *Incubator) VerifyIncubatorEmail() error {
	if !utils.ValidateEmail(inc.Email) {
		return errors.New("invalid email")
	}
	return nil
}

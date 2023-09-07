package models

import (
	"errors"
	"pi/internal/api/v1/utils"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CompanyID string `json:"companyId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (usr *User) VerifyUserEmail() error {
	if !utils.ValidateEmail(usr.Email) {
		return errors.New("invalid email")
	}
	return nil
}

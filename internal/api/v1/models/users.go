package models

import (
	"errors"
	"pi/internal/api/v1/utils"
)

var UsersProps []string = []string{"owner", "other"}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Type      string `json:"type"`
	CompanyID string `json:"companyId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (usr *User) VerifyUserType() error {
	for _, prop := range UsersProps {
		if prop == usr.Type {
			return nil
		}
	}
	usr.Type = "other"
	return nil
}

func (usr *User) VerifyUserEmail() error {
	if !utils.ValidateEmail(usr.Email) {
		return errors.New("invalid email")
	}
	return nil
}

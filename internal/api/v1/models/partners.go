package models

import "errors"

var PartnerTypes []string = []string{"admin", "other"}

type Partner struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Type      string `json:"type"`
	CompanyID string `json:"companyId"`
	AccountID string `json:"accountId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (p *Partner) VerifyType() error {
	for _, prop := range PartnerTypes {
		if p.Type == prop {
			return nil
		}
	}

	return errors.New("invalid partner type")
}

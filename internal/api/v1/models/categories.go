package models

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyID string `json:"companyId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

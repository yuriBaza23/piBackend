package models

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CmpID       string `json:"companyID"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

package models

type Tasks struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	InitialDate string `json:"initDate"`
	EndDate     string `json:"endDate"`
	Status      string `json:"status"`
	ProjectID   string `json:"projectId"`
	CompanyID   string `json:"companyId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

package models

var FinanceProps []string = []string{"expense", "revenue"}

type Finance struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Category  string `json:"category"`
	Value     int    `json:"value"`
	FinValue  string `json:"finValue,omitempty"`
	CompanyID string `json:"companyId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (fin *Finance) VerifyFinaceType() error {
	for _, prop := range FinanceProps {
		if prop == fin.Type {
			return nil
		}
	}
	fin.Type = "revenue"
	return nil
}

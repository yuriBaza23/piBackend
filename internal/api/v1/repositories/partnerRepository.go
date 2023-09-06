package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func InsertPartner(partner models.Partner) (id string, err error) {
	err = partner.VerifyType()
	if err != nil {
		return
	}

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO partners (id, name, email, type, companyId, accountId, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err = db.QueryRow(stmt, partner.ID, partner.Name, partner.Email, partner.Type, partner.CompanyID, partner.AccountID, time.Now(), time.Now()).Scan(&id)

	return
}

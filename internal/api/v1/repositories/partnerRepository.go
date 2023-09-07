package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeletePartner(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM partners WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertPartner(partner models.Partner) (id string, err error) {
	err = partner.VerifyType()
	if err != nil {
		return
	}

	err = partner.VerifyEmail()
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

func GetPartner(id string) (partner models.Partner, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM partners WHERE id=$1`

	err = conn.QueryRow(stmt, id).Scan(&partner.ID, &partner.Name, &partner.Type, &partner.Email, &partner.CompanyID, &partner.AccountID, &partner.CreatedAt, &partner.UpdatedAt)

	return
}

func GetAllPartners() (s []models.Partner, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT id, name, email, type, companyId, accountId, createdAt, updatedAt FROM partners`
	rows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for rows.Next() {
		var partner models.Partner

		err = rows.Scan(&partner.ID, &partner.Name, &partner.Email, &partner.Type, &partner.CompanyID, &partner.AccountID, &partner.CreatedAt, &partner.UpdatedAt)
		if err != nil {
			continue
		}

		s = append(s, partner)
	}

	return
}

func UpdatePartner(id string, partner models.Partner) (int64, error) {
	err := partner.VerifyType()
	if err != nil {
		return 0, err
	}

	err = partner.VerifyEmail()
	if err != nil {
		return 0, err
	}

	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE partners SET name=$1, email=$2, type=$3, accountId=$4, updatedAt=$5 WHERE id=$6`
	row, err := conn.Exec(stmt, partner.Name, partner.Email, partner.Type, partner.AccountID, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

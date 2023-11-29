package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeleteFinance(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM finances WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertFinance(fin models.Finance) (id string, err error) {
	err = fin.VerifyFinaceType()
	if err != nil {
		return
	}

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO finances (id, name, type, category, value, companyId) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = db.QueryRow(stmt, fin.ID, fin.Name, fin.Type, fin.Category, fin.Value, fin.CompanyID).Scan(&id)

	return
}

func GetFinance(id string) (fin models.Finance, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM finances WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&fin.ID, &fin.Name, &fin.Type, &fin.Category, &fin.Value, &fin.CompanyID, &fin.CreatedAt, &fin.UpdatedAt)

	return
}

func GetAllCompanyFinance(id string) (fin []models.Finance, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT id, name, type, category, value, companyId, createdAt, updatedAt FROM finances WHERE companyId=$1`
	finRows, err := conn.Query(stmt, id)
	if err != nil {
		return
	}

	for finRows.Next() {
		var f models.Finance

		err = finRows.Scan(&f.ID, &f.Name, &f.Type, &f.Category, &f.Value, &f.CompanyID, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			continue
		}

		fin = append(fin, f)
	}

	return
}

func UpdateFinance(id string, fin models.Finance) (int64, error) {
	err := fin.VerifyFinaceType()
	if err != nil {
		return 0, err
	}

	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE finances SET name=$1, type=$2, category=$3, value=$4, updatedAt=$5 WHERE id=$6`
	row, err := conn.Exec(stmt, fin.Name, fin.Type, fin.Category, fin.Value, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

package repositories

import (
	"pi/internal/api/v1/models"
)

func DeleteCompany(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM companies WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertCompany(cmp models.Company) (id string, err error) {
	err = cmp.VerifyCompanyEmail()
	if err != nil {
		return
	}

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO companies (id, name, email, cnpj, hubId) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = db.QueryRow(stmt, cmp.ID, cmp.Name, cmp.Email, cmp.CNPJ, cmp.HubID).Scan(&id)

	return
}

func GetCompany(id string) (cmp models.Company, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM companies WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&cmp.ID, &cmp.Name, &cmp.Email, &cmp.CNPJ, &cmp.HubID, &cmp.CreatedAt, &cmp.UpdatedAt)

	userIds := []string{}
	stmt = `SELECT userId FROM users_companies WHERE companyId=$1`
	err = conn.QueryRow(stmt, id).Scan(&userIds)

	return
}

func GetCompanyByCNPJ(cnpj string) (cmp models.Company, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM companies WHERE cnpj=$1`
	err = conn.QueryRow(stmt, cnpj).Scan(&cmp.ID, &cmp.Name, &cmp.Email, &cmp.CNPJ, &cmp.HubID, &cmp.CreatedAt, &cmp.UpdatedAt)

	userIds := []string{}
	stmt = `SELECT userId FROM users_companies WHERE companyId=$1`
	err = conn.QueryRow(stmt, cmp.ID).Scan(&userIds)

	return
}

package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeleteCompany(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM users_companies WHERE companyId=$1`
	_, err = conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	stmt = `DELETE FROM companies WHERE id=$1`
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

	stmt = `SELECT userId, type FROM users_companies WHERE companyId=$1`
	rows, err := conn.Query(stmt, id)
	if err != nil {
		return
	}

	for rows.Next() {
		var usr models.User
		var userId string

		err = rows.Scan(&userId, &usr.Type)
		if err != nil {
			continue
		}

		stmt = `SELECT id, email, name, isPreRegister, createdAt, updatedAt FROM users WHERE id=$1`
		err = conn.QueryRow(stmt, userId).Scan(&usr.ID, &usr.Email, &usr.Name, &usr.IsPreReg, &usr.CreatedAt, &usr.UpdatedAt)
		if err != nil {
			continue
		}
		usr.CompanyID = id
		cmp.Users = append(cmp.Users, usr)
	}

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

	stmt = `SELECT userId FROM users_companies WHERE companyId=$1`
	rows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for rows.Next() {
		var usr models.User
		var userId string

		err = rows.Scan(&userId)
		if err != nil {
			continue
		}

		stmt = `SELECT id, email, name, isPreRegister, createdAt, updatedAt FROM users WHERE id=$1`
		err = conn.QueryRow(stmt, userId).Scan(&usr.ID, &usr.Email, &usr.Name, &usr.IsPreReg, &usr.CreatedAt, &usr.UpdatedAt)
		if err != nil {
			continue
		}
		usr.CompanyID = cmp.ID
		cmp.Users = append(cmp.Users, usr)
	}

	return
}

func GetAllCompanies() (cmp []models.Company, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM companies`
	compRows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for compRows.Next() {
		var c models.Company

		err = compRows.Scan(&c.ID, &c.Name, &c.Email, &c.CNPJ, &c.HubID, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			continue
		}

		stmt = `SELECT userId, type FROM users_companies WHERE companyId=$1`
		rows, err2 := conn.Query(stmt, c.ID)
		if err2 != nil {
			return
		}

		for rows.Next() {
			var usr models.User
			var userId string

			err = rows.Scan(&userId, &usr.Type)
			if err != nil {
				continue
			}

			stmt = `SELECT id, email, name, isPreRegister, createdAt, updatedAt FROM users WHERE id=$1`
			err = conn.QueryRow(stmt, userId).Scan(&usr.ID, &usr.Email, &usr.Name, &usr.IsPreReg, &usr.CreatedAt, &usr.UpdatedAt)
			if err != nil {
				continue
			}
			usr.CompanyID = c.ID
			c.Users = append(c.Users, usr)
		}

		cmp = append(cmp, c)
	}

	return
}

func UpdateCompany(id string, cmp models.Company) (int64, error) {
	err := cmp.VerifyCompanyEmail()
	if err != nil {
		return 0, err
	}

	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE companies SET name=$1, email=$2, cnpj=$3, hubId=$4, updatedAt=$5 WHERE id=$6`
	row, err := conn.Exec(stmt, cmp.Name, cmp.Email, cmp.CNPJ, cmp.HubID, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

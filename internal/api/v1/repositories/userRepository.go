package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeleteIncubator(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM Incubators WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertIncubator(inc models.Incubator) (id string, err error) {
	err = inc.VerifyIncubatorEmail()
	if err != nil {
		return
	}

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO Incubators (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`

	err = db.QueryRow(stmt, inc.ID, inc.Name, inc.Email, inc.Password).Scan(&id)

	stmt = `INSERT INTO Incubators_companies (companyId, IncubatorId) VALUES ($1, $2) RETURNING id`

	err = db.QueryRow(stmt, inc.CompanyID, inc.ID).Scan(&id)

	return inc.ID, err
}

func GetIncubator(id string) (inc models.Incubator, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM Incubators WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&inc.ID, &inc.Name, &inc.Email, &inc.Password, &inc.CreatedAt, &inc.UpdatedAt)

	stmt = `SELECT companyId FROM Incubators_companies WHERE IncubatorId=$1`
	err = conn.QueryRow(stmt, id).Scan(&inc.CompanyID)

	return
}

func GetAllIncubators() (incArray []models.Incubator, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT id, name, email, createdAt, updatedAt FROM Incubators`
	rows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for rows.Next() {
		var inc models.Incubator

		err = rows.Scan(&inc.ID, &inc.Name, &inc.Email, &inc.CreatedAt, &inc.UpdatedAt)
		if err != nil {
			continue
		}

		stmt = `SELECT companyId FROM Incubators_companies WHERE IncubatorId=$1`
		err = conn.QueryRow(stmt, inc.ID).Scan(&inc.CompanyID)
		if err != nil {
			continue
		}

		incArray = append(incArray, inc)
	}

	return
}

func UpdateIncubator(id string, inc models.Incubator) (int64, error) {
	err := inc.VerifyIncubatorEmail()
	if err != nil {
		return 0, err
	}

	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE Incubators SET name=$1, email=$2, password=$3, updatedAt=$4 WHERE id=$5`
	_, err = conn.Exec(stmt, inc.Name, inc.Email, inc.Password, time.Now(), id)
	if err != nil {
		return 0, err
	}

	stmt = `UPDATE Incubators_companies SET companyId=$1 WHERE IncubatorId=$2`
	row, err := conn.Exec(stmt, inc.CompanyID, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

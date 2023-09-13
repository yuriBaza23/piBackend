package repositories

import (
	"pi/cmd/internal/api/v1/models"
	"time"
)

func DeleteIncubator(id string) (int64, error) {
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

	stmt := `INSERT INTO incubators (id, name, email) VALUES ($1, $2, $3) RETURNING id`

	err = db.QueryRow(stmt, inc.ID, inc.Name, inc.Email).Scan(&id)

	return
}

func GetIncubator(id string) (inc models.Incubator, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM incubators WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&inc.ID, &inc.Name, &inc.Email, &inc.CreatedAt, &inc.UpdatedAt)

	return
}

func GetAllIncubators() (inc []models.Incubator, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM incubators`
	incRows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for incRows.Next() {
		var i models.Incubator

		err = incRows.Scan(&i.ID, &i.Name, &i.Email, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			continue
		}

		inc = append(inc, i)
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

	stmt := `UPDATE incubators SET name=$1, email=$2, updatedAt=$3 WHERE id=$4`
	row, err := conn.Exec(stmt, inc.Name, inc.Email, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

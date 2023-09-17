package repositories

import (
	"pi/cmd/internal/api/v1/models"
	"time"
)

func DeleteWarning(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM warnings WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertWarning(war models.Warning) (id string, err error) {

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO warnings (id, title, content, companyID) VALUES ($1, $2, $3, $4) RETURNING id`

	err = db.QueryRow(stmt, war.ID, war.Title, war.Content, war.CmpId).Scan(&id)

	return
}

func GetWarning(id string) (war models.Warning, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM warnings WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&war.ID, &war.Title, &war.Content, &war.CmpId, &war.CreatedAt, &war.UpdatedAt)

	return
}

func GetAllWarnings() (war []models.Warning, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM warnings`
	warRows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for warRows.Next() {
		var w models.Warning

		err = warRows.Scan(&w.ID, &w.Title, &w.Content, &w.CmpId, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			continue
		}

		war = append(war, w)
	}

	return
}

func UpdateWarning(id string, war models.Warning) (int64, error) {

	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE warnings SET title=$1, content=$2, updatedAt=$3 WHERE id=$4`
	row, err := conn.Exec(stmt, war.Title, war.Content, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeleteCategory(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM categories WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertCategory(cat models.Category) (id string, err error) {

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO categories (id, name, type, companyId) VALUES ($1, $2, $3, $4) RETURNING id`

	err = db.QueryRow(stmt, cat.ID, cat.Name, cat.Type, cat.CompanyID).Scan(&id)

	return
}

func GetCategory(id string) (cat models.Category, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM categories WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&cat.ID, &cat.Name, &cat.Type, &cat.CompanyID, &cat.CreatedAt, &cat.UpdatedAt)

	return
}

func GetAllCategories(id string) (cat []models.Category, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT id, name, type, companyId, createdAt, updatedAt FROM categories WHERE companyId=$1`
	catRows, err := conn.Query(stmt, id)
	if err != nil {
		return
	}

	for catRows.Next() {
		var c models.Category

		err = catRows.Scan(&c.ID, &c.Name, &c.Type, &c.CompanyID, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			continue
		}

		cat = append(cat, c)
	}

	return
}

func UpdateCategory(id string, cat models.Category) (int64, error) {

	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE categories SET name=$1, type=$2, updatedAt=$3 WHERE id=$4`
	row, err := conn.Exec(stmt, cat.Name, cat.Type, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

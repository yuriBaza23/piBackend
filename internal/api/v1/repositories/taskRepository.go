package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeleteTask(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM tasks WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertTask(tsk models.Tasks) (id string, err error) {
	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO tasks (id, title, description, initialDate, endDate, status, projectId, companyId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err = db.QueryRow(stmt, tsk.ID, tsk.Title, tsk.Description, tsk.InitialDate, tsk.EndDate, tsk.Status, tsk.ProjectID, tsk.CompanyID).Scan(&id)

	return
}

func GetTask(id string) (tsk models.Tasks, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM task WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&tsk.ID, &tsk.Title, &tsk.Description, &tsk.InitialDate, &tsk.EndDate, &tsk.Status, &tsk.ProjectID, &tsk.CompanyID, &tsk.CreatedAt, &tsk.UpdatedAt)

	return
}

func GetAllProjectTasks(id string) (tsk []models.Tasks, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT id, title, description, initialDate, endDate, status, projectId, companyId, createdAt, updatedAt FROM tasks WHERE projectId=$1`
	tskRows, err := conn.Query(stmt, id)
	if err != nil {
		return
	}

	for tskRows.Next() {
		var t models.Tasks

		err = tskRows.Scan(&t.ID, &t.Title, &t.Description, &t.InitialDate, &t.EndDate, &t.Status, &t.ProjectID, &t.CompanyID, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			continue
		}

		tsk = append(tsk, t)
	}

	return
}

func UpdateTask(id string, tsk models.Tasks) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE tasks SET title=$1, description=$2, initialDate=$3, endDate=$4, status=$5, updatedAt=$6 WHERE id=$7`
	row, err := conn.Exec(stmt, tsk.Title, tsk.Description, tsk.InitialDate, tsk.EndDate, tsk.Status, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeleteProject(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM projects WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertProject(prj models.Project) (id string, err error) {

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO projects (id, name, companyID, description) VALUES ($1, $2, $3, $4) RETURNING id`

	err = db.QueryRow(stmt, prj.ID, prj.Name, prj.CmpID, prj.Description).Scan(&id)

	return
}

func GetProject(id string) (prj models.Project, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM projects WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&prj.ID, &prj.Name, &prj.CmpID, &prj.Description, &prj.CreatedAt, &prj.UpdatedAt)

	return
}

func GetProjectByCompany(companyID string) (prj models.Project, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM projects WHERE companyID=$1`
	err = conn.QueryRow(stmt, companyID).Scan(&prj.ID, &prj.Name, &prj.CmpID, &prj.CreatedAt, &prj.UpdatedAt)

	return
}

func GetAllProjects() (prj []models.Project, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT id, name, companyID, description, createdAt, updatedAt FROM projects`
	prjRows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for prjRows.Next() {
		var p models.Project

		err = prjRows.Scan(&p.ID, &p.Name, &p.CmpID, &p.Description, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			continue
		}

		prj = append(prj, p)
	}

	return
}

func UpdateProject(id string, prj models.Project) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE projects SET name=$1, companyID=$2, description=$3, updatedAt=$4 WHERE id=$5`
	row, err := conn.Exec(stmt, prj.Name, prj.CmpID, prj.Description, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

package repositories

import (
	"pi/internal/api/v1/models"
	"time"
)

func DeleteUser(id string) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM users WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertUser(usr models.User) (id string, err error) {
	err = usr.VerifyUserEmail()
	if err != nil {
		return
	}

	db, err := OpenConnection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`

	err = db.QueryRow(stmt, usr.ID, usr.Name, usr.Email, usr.Password).Scan(&id)

	stmt = `INSERT INTO users_companies (companyId, userId, type) VALUES ($1, $2, $3) RETURNING id`

	err = db.QueryRow(stmt, usr.CompanyID, usr.ID, usr.Type).Scan(&id)

	return usr.ID, err
}

func GetUser(id string) (usr models.User, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM users WHERE id=$1`
	err = conn.QueryRow(stmt, id).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)

	stmt = `SELECT companyId, type FROM users_companies WHERE userId=$1`
	err = conn.QueryRow(stmt, id).Scan(&usr.CompanyID, &usr.Type)

	return
}

func GetAllUsers() (usrArray []models.User, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT id, name, email, createdAt, updatedAt FROM users`
	rows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for rows.Next() {
		var usr models.User

		err = rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.CreatedAt, &usr.UpdatedAt)
		if err != nil {
			continue
		}

		stmt = `SELECT companyId, type FROM users_companies WHERE userId=$1`
		err = conn.QueryRow(stmt, usr.ID).Scan(&usr.CompanyID, &usr.Type)
		if err != nil {
			continue
		}

		usrArray = append(usrArray, usr)
	}

	return
}

func UpdateUser(id string, usr models.User) (int64, error) {
	err := usr.VerifyUserEmail()
	if err != nil {
		return 0, err
	}

	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt := `UPDATE users SET name=$1, email=$2, password=$3, updatedAt=$4 WHERE id=$5`
	_, err = conn.Exec(stmt, usr.Name, usr.Email, usr.Password, time.Now(), id)
	if err != nil {
		return 0, err
	}

	stmt = `UPDATE users_companies SET companyId=$1, type=$2 WHERE userId=$3`
	row, err := conn.Exec(stmt, usr.CompanyID, usr.Type, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func GetUserByEmail(email string) (usr models.User, err error) {
	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM users WHERE email=$1`
	err = conn.QueryRow(stmt, email).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)

	stmt = `SELECT companyId, type FROM users_companies WHERE userId=$1`
	err = conn.QueryRow(stmt, usr.ID).Scan(&usr.CompanyID, &usr.Type)

	return
}

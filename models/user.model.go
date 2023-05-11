package models

import (
	"fmt"
	"time"

	"github.com/rfauzi44/online-course-api/db"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,min=6"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted,omitempty"`
	Name      string    `json:"name" validate:"required"`
}

func ReadAllUser() (*[]User, error) {
	var data User
	var dataArray []User

	conn := db.Connect()

	sqlStatement := `SELECT id, email, role, created_at, updated_at, name
	FROM users 
	WHERE is_deleted = false;`

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Email, &data.Role, &data.CreatedAt, &data.UpdatedAt, &data.Name)
		if err != nil {
			return nil, err
		}

		dataArray = append(dataArray, data)
	}
	return &dataArray, nil

}

func DeleteUser(email string) (interface{}, error) {

	conn := db.Connect()

	sqlStatement := `UPDATE users SET is_deleted = ? WHERE email = ?`

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(true, email)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("user with email %s not found", email)
	}

	data := map[string]int64{
		"rows_affected": rowsAffected,
	}

	return data, nil
}

func ChangeRole(email string) (interface{}, error) {

	conn := db.Connect()

	sqlStatement := `UPDATE users SET role = ? WHERE email = ?`

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec("admin", email)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("user with email %s not found", email)
	}

	data := map[string]int64{
		"rows_affected": rowsAffected,
	}

	return data, nil
}

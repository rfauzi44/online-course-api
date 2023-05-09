package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rfauzi44/online-course-api/db"
	"github.com/rfauzi44/online-course-api/libs"
)

func Register(email, password string) (interface{}, error) {
	uuid := uuid.New()
	id := uuid.String()

	conn := db.Connect()

	sqlStatement := "INSERT users (id, email, password, role, created_at, updated_at, is_deleted) VALUES (?, ?, ?, ?, ?, ?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id, email, password, "user", time.Now(), time.Now(), false)
	if err != nil {
		return nil, err
	}

	data := map[string]string{
		"last_inserted_id": id,
	}

	return data, nil
}

func Login(email, password string) (*User, error) {

	var data User
	conn := db.Connect()

	sqlStatement := "SELECT * FROM users WHERE email = ?"
	err := conn.QueryRow(sqlStatement, email).Scan(
		&data.ID, &data.Email, &data.Password, &data.Role, &data.CreatedAt, &data.UpdatedAt, &data.IsDeleted,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not registered")
	}

	if err != nil {
		return nil, err
	}

	match, _ := libs.ComparePassword(password, data.Password)
	if !match {
		return nil, errors.New("wrong password")
	}

	return &data, nil
}

package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rfauzi44/online-course-api/db"
	"github.com/rfauzi44/online-course-api/libs"
)

func Register(user User) (*User, error) {
	id := uuid.New().String()

	conn := db.Connect()

	sqlStatement := `SELECT email FROM users WHERE email = ?`
	err := conn.QueryRow(sqlStatement, user.Email).Scan(
		&user.Email,
	)

	if err != sql.ErrNoRows {
		return nil, errors.New("email already registered")
	}

	sqlStatement = `INSERT users (id, email, password, role, created_at, updated_at, is_deleted, name) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id, user.Email, user.Password, "user", time.Now(), time.Now(), false, user.Name)
	if err != nil {
		return nil, err
	}

	sqlStatement = `SELECT id, email, role, created_at, updated_at, name FROM users WHERE id = ?`

	var registeredUser User
	err = conn.QueryRow(sqlStatement, id).Scan(
		&registeredUser.ID,
		&registeredUser.Email,
		&registeredUser.Role,
		&registeredUser.CreatedAt,
		&registeredUser.UpdatedAt,
		&registeredUser.Name,
	)
	if err != nil {
		return nil, err
	}

	return &registeredUser, nil
}

func Login(email, password string) (*User, error) {

	var user User
	conn := db.Connect()

	sqlStatement := `SELECT id, email, password, role, created_at, updated_at, name FROM users WHERE is_deleted = ? AND email = ?`
	err := conn.QueryRow(sqlStatement, false, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Name,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("email not registered")
	}

	if err != nil {
		return nil, err
	}

	valid, _ := libs.ValidatePassword(password, user.Password)
	if !valid {
		return nil, errors.New("wrong password")
	}

	user.Password = ""

	return &user, nil
}

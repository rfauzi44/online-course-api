package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rfauzi44/online-course-api/db"
	"github.com/rfauzi44/online-course-api/libs"
)

type Course struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" form:"title" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	Price       int       `json:"price" form:"price" validate:"gte=0"`
	Image       string    `json:"image" form:"image"`
	Category    string    `json:"category" form:"category" validate:"required"`
	AuthorID    string    `json:"author_id,omitempty" form:"author_id"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ImageID     string    `json:"image_id,omitempty"`
}

func CreateCourse(course Course) (*Course, error) {

	course.ID = uuid.New().String()
	conn := db.Connect()

	sqlStatement := `INSERT courses (id, title, description, price, image, category, author_id, created_at, updated_at, image_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(course.ID, course.Title, course.Description, course.Price, course.Image, course.Category, course.AuthorID, time.Now(), time.Now(), course.ImageID)
	if err != nil {
		return nil, err
	}

	sqlStatement = `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id WHERE c.id = ?`
	row := conn.QueryRow(sqlStatement, course.ID)

	var data Course

	err = row.Scan(&data.ID, &data.Title, &data.Description, &data.Price, &data.Image, &data.Category, &data.Author, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func ReadAllCourse() ([]Course, error) {
	var data Course
	var dataArray []Course

	conn := db.Connect()

	sqlStatement := `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id`

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Title, &data.Description, &data.Price, &data.Image, &data.Category, &data.Author, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}

		dataArray = append(dataArray, data)
	}
	return dataArray, nil

}

func UpdateCourse(course Course) (*Course, error) {

	conn := db.Connect()

	sqlStatement := `UPDATE courses SET title = ?, description = ?, price = ?, image = ?, category = ?, updated_at = ?, image_id = ? WHERE id = ?`

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(course.Title, course.Description, course.Price, course.Image, course.Category, time.Now(), course.ImageID, course.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("course with id %s not found", course.ID)

	}

	sqlStatement = `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id WHERE c.id = ?`
	row := conn.QueryRow(sqlStatement, course.ID)

	var data Course

	err = row.Scan(&data.ID, &data.Title, &data.Description, &data.Price, &data.Image, &data.Category, &data.Author, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func DeleteCourse(id string) (interface{}, error) {

	conn := db.Connect()

	sqlStatement := `SELECT image_id
	FROM courses 
	WHERE id = ?`
	row := conn.QueryRow(sqlStatement, id)

	var course Course

	err := row.Scan(&course.ImageID)
	if err != nil {
		return nil, err
	}

	sqlStatement = `DELETE FROM courses WHERE id = ?`

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("course with id %s not found", id)

	}

	_, err = libs.DeleteImage(course.ImageID)
	if err != nil {
		return nil, err
	}

	data := map[string]int64{
		"rows_affected": rowsAffected,
	}

	return data, nil
}

func GetCourseById(id string) (*Course, error) {
	var data Course

	conn := db.Connect()

	sqlStatement := `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id WHERE c.id = ?`

	row := conn.QueryRow(sqlStatement, id)

	err := row.Scan(&data.ID, &data.Title, &data.Description, &data.Price, &data.Image, &data.Category, &data.Author, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("course with id %s not found", id)
		}
		return nil, err
	}

	return &data, nil

}

func SearchCourse(query string) ([]Course, error) {
	var data Course
	var dataArray []Course

	conn := db.Connect()

	sqlStatement := `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id WHERE c.title LIKE ? OR c.description LIKE ? OR c.category LIKE ? OR u.name LIKE ?`

	rows, err := conn.Query(sqlStatement, "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Title, &data.Description, &data.Price, &data.Image, &data.Category, &data.Author, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}

		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}

func SortCourse(query string) ([]Course, error) {
	var data Course
	var dataArray []Course

	conn := db.Connect()

	var sqlStatement string
	switch query {
	case "high":
		sqlStatement = `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id
	WHERE c.price > 0
	ORDER BY c.price DESC`
	case "low":
		sqlStatement = `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id 
	WHERE c.price > 0
	ORDER BY c.price ASC`
	case "free":
		sqlStatement = `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id WHERE c.price = 0`
	default:
		sqlStatement = `SELECT c.id, c.title, c.description, c.price, c.image, c.category, u.name, c.created_at, c.updated_at
	FROM courses c
	JOIN users u ON c.author_id = u.id`
	}

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Title, &data.Description, &data.Price, &data.Image, &data.Category, &data.Author, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}

		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}

func GetCategory() ([]string, error) {
	var data string
	var dataArray []string

	conn := db.Connect()

	rows, err := conn.Query(`SELECT DISTINCT category FROM courses`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&data)
		if err != nil {
			return nil, err
		}
		dataArray = append(dataArray, data)
	}

	return dataArray, nil
}

type PopularCategory struct {
	Category    string `json:"category"`
	TotalCourse int    `json:"total_course"`
}

func GetPopularCategory() ([]PopularCategory, error) {
	var data PopularCategory
	var dataArray []PopularCategory

	conn := db.Connect()

	rows, err := conn.Query(`SELECT category, COUNT(*) as count FROM courses GROUP BY category ORDER BY count DESC LIMIT 3`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&data.Category, &data.TotalCourse)
		if err != nil {
			return nil, err
		}
		dataArray = append(dataArray, data)
	}

	return dataArray, nil
}

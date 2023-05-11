package models

import "github.com/rfauzi44/online-course-api/db"

type Stats struct {
	TotalUser       int `json:"total_user"`
	TotalCourse     int `json:"total_course"`
	TotalFreeCourse int `json:"total_free_course"`
}

func GetStats() (*Stats, error) {
	var data Stats

	conn := db.Connect()

	err := conn.QueryRow(`SELECT COUNT(*) FROM users WHERE is_deleted = false`).Scan(&data.TotalUser)
	if err != nil {
		return nil, err
	}

	err = conn.QueryRow(`SELECT COUNT(*) FROM courses`).Scan(&data.TotalCourse)
	if err != nil {
		return nil, err
	}

	err = conn.QueryRow(`SELECT COUNT(*) FROM courses WHERE price = 0`).Scan(&data.TotalFreeCourse)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {

	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DBNAME := os.Getenv("MYSQL_DBNAME")

	connectionString := MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DBNAME + "?charset=utf8mb4&parseTime=True"

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("database error in connection string")
	}

	err = db.Ping()
	if err != nil {
		panic("database ping unsuccessful")
	}

}

func Connect() *sql.DB {
	return db

}

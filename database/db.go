package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/gorilla?parseTime=true")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

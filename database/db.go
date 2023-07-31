package database

import (
	"database/sql"
	"fmt"
	"time"

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

	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		return nil, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var (
		queried_id        int
		queried_username  string
		queried_password  string
		queried_createdAt time.Time
	)

	query = `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err = db.QueryRow(query, 1).Scan(&queried_id, &queried_username, &queried_password, &queried_createdAt)
	if err != nil {
		return nil, err
	}
	fmt.Println("userId: ", userId)
	fmt.Println("userName: ", queried_username)
	fmt.Println("password: ", queried_password)
	fmt.Println("createdAt: ", queried_createdAt)
	return db, nil
}

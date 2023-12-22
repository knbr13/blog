package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/iline")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUserByID(db *sql.DB, userID int) (*User, error) {
	var user User
	err := db.QueryRow("select * from users where id = ?", userID).Scan(&user.ID, &user.Name)
	return &user, err
}

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackskj/carta"
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

	db.SetMaxOpenConns(1)
}

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

func GetUsers(db *sql.DB) ([]User, error) {
	var users []User
	rows, err := db.Query("select * from users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(db *sql.DB, userID int) (*User, error) {
	var user User
	err := db.QueryRow("select * from users where id = ?", userID).Scan(&user.ID, &user.Name)
	return &user, err
}

func getUserWithAccounts(db *sql.DB, userID int) (*User, error) {
	var user User
	rows, err := db.Query("select * from users left join accounts where id = ?", userID)
	if err != nil {
		return nil, err
	}

	err = carta.Map(rows, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

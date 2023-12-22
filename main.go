package main

import (
	"database/sql"
	"fmt"
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

	fmt.Println("fetching user...")

	user1, err := getUserByID(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user1: %+v\n", user1)

	fmt.Println("fetching user...")

	user2, err := getUserByID(db, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user2: %+v\n", user2)

	fmt.Println("fetching users...")

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	for i, user := range users {
		fmt.Println("user:", i+1, ":", user)
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(db *sql.DB) ([]User, error) {
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

func getUserByID(db *sql.DB, userID int) (*User, error) {
	var user User
	err := db.QueryRow("select * from users where id = ?", userID).Scan(&user.ID, &user.Name)
	return &user, err
}

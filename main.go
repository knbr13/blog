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
}

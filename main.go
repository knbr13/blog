package main

import (
	"fmt"
	"go-arg/database"
)

func main() {
	_, err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to database ", err)
		return
	}
	// r := mux.NewRouter()

	// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]

	// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// })

	// http.ListenAndServe(":8080", r)
}

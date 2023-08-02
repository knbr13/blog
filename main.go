package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode/", func(w http.ResponseWriter, r *http.Request) {
		var user *User
		json.NewDecoder(r.Body).Decode(user)
		fmt.Fprintf(w, "%s %s is %d years old!", user.FirstName, user.LastName, user.Age)
	})
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
	}).Methods(http.MethodGet)

	router.HandleFunc("/post/{age}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		age := vars["age"]
		ageInt, err := strconv.Atoi(age)
		if err != nil || ageInt < 1 {
			http.Error(w, "Error parsing age, please make sure it is a valid number and greater than 0", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "your age is %d\n", ageInt)
	})

	log.Println("starting the server...")
	log.Fatalln(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
	}).Methods(http.MethodGet)

	log.Println("starting the server...")
	log.Fatalln(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

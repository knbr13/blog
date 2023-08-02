package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", chain(sayHello, logging, method("GET")))
	http.ListenAndServe(":8080", nil)
}

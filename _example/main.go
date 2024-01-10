package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

func main() {
	http.Handle("/", httprate.LimitByIP(10, time.Minute)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world!\n")
	})))
	http.ListenAndServe(":8080", nil)
}

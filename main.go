package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() { log.Println(time.Since(start)) }()
		f(w, r)
	}
}

func method(name string, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != name {
			return
		}
		f(w, r)
	}
}

func chain(f http.HandlerFunc, mds ...Middleware) http.HandlerFunc {
	for _, md := range mds {
		f = md(f)
	}
	return f
}

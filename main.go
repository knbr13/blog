package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			hf(w, r)
		}
	}
}

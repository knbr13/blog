package main

import (
	"net/http"
)

func main() {

}

type Middleware func(http.HandlerFunc) http.HandlerFunc

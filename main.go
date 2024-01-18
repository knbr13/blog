package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("hello.txt", []byte("hello world from hello.txt!"), 0655)
	if err != nil {
		log.Fatal(err)
	}
}

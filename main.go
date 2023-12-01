package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	x := "GET / HTTP/1.1\r\nAccept:*/*\r\nContent-Type: text/plain\r\n\r\nhello world"
	f, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, x)
}

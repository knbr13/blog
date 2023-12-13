package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	os.Stdout = f
	fmt.Fprintln(os.Stdout, "hello world!")
	fmt.Println("this is me!")
	println()
}

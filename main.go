package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("size of file %q: %d\n", f.Name(), stat.Size())
}

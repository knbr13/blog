package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("secret.file.txt")
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	fmt.Println("buf len:", buf.Len())
	fmt.Println("buf cap:", buf.Cap())
	n, err := io.Copy(&buf, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("n:", n)

	fmt.Println("buf len:", buf.Len())
	fmt.Println("buf cap:", buf.Cap())
	fmt.Println(buf.String())
}

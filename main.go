package main

import "fmt"

func main() {
	hello := []byte("hello")
	helloCopy := make([]byte, 5)
	copy(helloCopy, hello)
	fmt.Println(string(helloCopy))
}

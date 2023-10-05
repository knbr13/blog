package main

import "fmt"

func main() {
	s := New[int]()
	s.Push(10)
	s.Push(30)
	s.Push(50)
	s.Push(70)
	s.Push(90)
	s.Pop()
	fmt.Println()
	fmt.Println(s)
}

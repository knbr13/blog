package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello_world_how_are_you_there?"
	ss := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_'
	})

	for _, v := range ss {
		fmt.Println(v)
	}
}

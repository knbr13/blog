package main

import (
	"fmt"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	s1 := "git"
	s2 := "go-cleanhttp"
	rank := fuzzy.LevenshteinDistance(s1, s2)
	fmt.Println(rank)
	fmt.Println((rank/(rank+1) + 1) * 25)
}

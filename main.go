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
	fmt.Println(len("vides an easy way to capture http related metrics (i.e. response time, bytes written, and http status code) from your application's http.Handlers. description: HTTP routing for Go 1.7 description: A delightfully tiny but powerful HTTP router for Go web applications description: A flexible and comprehensive Kotlin library for encapsulating the resul")) // maximum description length on github
	fmt.Println(len("binwiederhier/ntfy  https://github.com/binwiederhier/ntfy  Send push notifications t...  15205  25"))
}

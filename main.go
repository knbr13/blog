package main

import "fmt"

func dedupLettersAndLexicographicOrder(s string) string {
	index := make([]bool, 26)
	for _, r := range s {
		if i := r - 'a'; !index[i] {
			index[i] = true
		}
	}

	res := make([]rune, 0, len(index))
	for i, ok := range index {
		if ok {
			res = append(res, rune(i+'a'))
		}
	}

	return string(res)
}

func main() {
	rubric := map[string]string{
		"aaaaaaabdeeeecfff": "abcdef",
		"bbbefffghaaaaaaa":  "abefgh",
	}

	for input, want := range rubric {
		got := dedupLettersAndLexicographicOrder(input)
		if got != want {
			fmt.Printf("%q =>\n\tGot:  %q\n\tWant: %q", input, got, want)
		}
	}
}

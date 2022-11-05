package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("aAbc"))
	fmt.Println(isUnique("aabc"))
}

func isUnique(s string) bool {
	set := map[rune]struct{}{}
	for _, c := range strings.ToLower(s) {

		if _, ok := set[c]; !ok {
			set[c] = struct{}{}
		} else {
			return false
		}

	}
	return true
}

/*Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false*/

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

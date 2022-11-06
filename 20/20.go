package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reversedWords("snow dog sun"))

}

func reversedWords(s string) string {
	splitedStr := strings.Split(s, " ")

	for i, j := 0, len(splitedStr)-1; i < j; i, j = i+1, j-1 {

		splitedStr[i], splitedStr[j] = splitedStr[j], splitedStr[i]
	}

	return strings.Join(splitedStr, " ")
}

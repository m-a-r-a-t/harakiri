package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(reverseStr("DBFDвабсд"))
}

func reverseStr(s string) string {
	runes := []rune(s)
	var reversedStr strings.Builder
	strLength := len(runes)

	for i := 0; i < strLength; i++ {
		reversedStr.WriteRune(runes[strLength-1-i])
	}

	return reversedStr.String()

}

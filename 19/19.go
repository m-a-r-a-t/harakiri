/*Разработать программу, которая переворачивает
подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(reverseStr("DBFDвабсд"))
}

func reverseStr(s string) string {
	runes := []rune(s) // переводим в руны так как могут быть символы unicode
	var reversedStr strings.Builder
	strLength := len(runes)

	for i := 0; i < strLength; i++ {
		reversedStr.WriteRune(runes[strLength-1-i])
	}

	return reversedStr.String()

}

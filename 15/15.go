package main

import (
	"fmt"
)

var justString string

func main() {
	fmt.Println("а"[1:])
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	ind := 100
	fmt.Println(len(v))
	if ind <= len(runes)-1 {
		justString := runes[:100]
		fmt.Println("String:", string(justString))

	} else {
		fmt.Println("Граница слайса больше чем кол-во символов в строке")
	}

}

func createHugeString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			s += "Ж"
		} else {
			s += "B"
		}
	}
	return s
}

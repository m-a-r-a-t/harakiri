package main

import "fmt"

func main() {
	a := 1
	b := 2

	a, b = b, a
	fmt.Println(a, b)

	a, b = 1, 2

	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)

}

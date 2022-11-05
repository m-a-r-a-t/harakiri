package main

import "fmt"

func main() {

	checkType(5)
	checkType("abc")
	checkType(true)
	checkType(make(chan int))

}

func checkType(param interface{}) {

	switch param.(type) {
	case int:
		fmt.Println("Is int64")
	case string:
		fmt.Println("Is string")
	case bool:
		fmt.Println("Is bool")
	case chan any:
		fmt.Println("Is chan")
	default:
		fmt.Println("unexpected type")

	}

}

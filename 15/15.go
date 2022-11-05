package main

import "fmt"

var justString string

// func someFunc() {
// 	v := createHugeString(int64(1) << 10)
// 	justString = v[:100]
// }

func main() {
	fmt.Println(int32(1) << 10)
	fmt.Println(int64(1) << 60)

	// someFunc()
}

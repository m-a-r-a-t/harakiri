package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 4

	fmt.Println(slice)

	slice = append(slice[0:i], slice[i+1:]...)
	fmt.Println(slice)
}

/*Удалить i-ый элемент из слайса.*/

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 4

	fmt.Println(slice)

	slice = append(slice[0:i], slice[i+1:]...) // если важен порядок
	fmt.Println(slice)

	slice = remove(slice, 0) // если порядок не важен
	fmt.Println(slice)
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1] // на место i ставим последний
	return s[:len(s)-1]
}

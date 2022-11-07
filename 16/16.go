/*Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/

package main

import "fmt"

func main() {
	intArr := []int{5, 12, 6, 71, 33, 3, 1, 2, 5}

	fmt.Println("Not sorted", intArr)

	fmt.Println("Sorted    ", quickSort(intArr))
}

func quickSort(data []int) []int {
	if len(data) == 0 {
		return []int{}
	}
	mid := data[len(data)/2]
	equals := []int{}
	over := []int{}
	less := []int{}

	for _, v := range data {
		if v == mid {
			equals = append(equals, v)
		} else if v > mid {
			over = append(over, v)
		} else {
			less = append(less, v)

		}
	}

	return append(append(quickSort(less), equals...), quickSort(over)...)

}

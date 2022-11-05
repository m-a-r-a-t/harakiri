package main

import (
	"fmt"
)

func main() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binarySearch(intArr, 9))
}

func binarySearch(data []int, value int) int {

	left := 0
	right := len(data) - 1

	for left <= right {
		mid := (left + right) / 2
		if data[mid] == value {
			return mid
		} else if data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

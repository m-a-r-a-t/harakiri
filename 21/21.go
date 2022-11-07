/*Реализовать паттерн «адаптер» на любом примере.*/

package main

import "fmt"

type ISolver interface {
	CalcTwoNums(a, b int) int
}

type Calculator struct {
}

func (c *Calculator) Sum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

type Adapter struct {
	Calculator
}

func (ad *Adapter) CalcTwoNums(a, b int) int {

	return ad.Sum(a, b)
}


func main() {

	var solver ISolver = &Adapter{}

	result := solver.CalcTwoNums(1, 2)

	fmt.Println(result)
}

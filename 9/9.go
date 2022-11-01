package main

import "fmt"

func main() {
	exitChan := make(chan int)
	numbersArr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersFromArrChan := make(chan int64)

	squareChan := make(chan int64)

	go func() {
		defer close(squareChan)
		for n := range numbersFromArrChan {
			squareChan <- n * n
		}
	}()

	go func() {
		for square := range squareChan {
			fmt.Println(square)
		}
		exitChan <- 0
	}()

	for _, val := range numbersArr {
		numbersFromArrChan <- val
	}

	close(numbersFromArrChan)

	<-exitChan
}

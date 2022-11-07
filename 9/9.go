/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
 во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.*/

package main

import "fmt"

func main() {
	exitChan := make(chan int)
	numbersArr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersFromArrChan := make(chan int64)

	squareChan := make(chan int64)

	go func() {
		defer close(squareChan) // после завершения канала numbersFromArrChan закрываем squareChan
		for n := range numbersFromArrChan {
			squareChan <- n * n
		}
	}()

	go func() {
		for square := range squareChan {
			fmt.Println(square)
		}
		exitChan <- 0 // после завершения squareChan отправится сигнал о завершении
	}()

	for _, val := range numbersArr {
		numbersFromArrChan <- val
	}

	close(numbersFromArrChan)

	<-exitChan
}

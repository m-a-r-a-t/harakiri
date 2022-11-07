/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.*/


package main

import (
	"fmt"
	"sync"
)

func main() {
	var wgSquare sync.WaitGroup
	numsArray := []int64{2, 4, 6, 8, 10}
	squareChan := make(chan int64)
	sumChan := make(chan int64)

	go func() { // горутина суммирования значений
		var sum int64 = 0
		for num := range squareChan {

			sum += num
		}
		sumChan <- sum
	}()

	// В цикле конкуретно запускаем функцию расчета для каждого элемента массива
	for _, n := range numsArray {
		wgSquare.Add(1) // Добавляем +1 к счетчику,
		go squareOfNumber(n, &wgSquare, squareChan)

	}

	wgSquare.Wait() // Блокировка до тех пор пока не выполняться все горутины
	close(squareChan) // Закрываем канал, чтобы завершилась горутина суммирования

	fmt.Println("Sum:", <-sumChan)
}

// Функция расчета  квадрата и вывода в stdout
func squareOfNumber(num int64, wg *sync.WaitGroup, squareChan chan<- int64) {
	defer wg.Done() // После выполнения счетчик waitgroup уменьшится
	squareChan <- num * num
}

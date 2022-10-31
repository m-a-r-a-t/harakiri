package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numsArray := []int64{2, 4, 6, 8, 10}

	// В цикле конкуретно запускаем функцию расчета для каждого элемента массива
	for _, n := range numsArray {
		wg.Add(1) // Добавляем +1 к счетчику,
		go squareOfNumber(n, &wg)

	}

	wg.Wait() // Блокировка до тех пор пока не выполняться все горутины
}

// Функция расчета  квадрата и вывода в stdout
func squareOfNumber(num int64, wg *sync.WaitGroup) {
	defer wg.Done() // После выполнения счетчик waitgroup уменьшится
	fmt.Println(fmt.Sprintf("square of  %d : %d", num, num*num))
}

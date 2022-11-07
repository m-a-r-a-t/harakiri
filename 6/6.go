/*Реализовать все возможные способы остановки выполнения горутины.*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stopChan := make(chan int)
	isStoped := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				fmt.Println("Завершение работы канала 1")
				return
			default:
				fmt.Println("Горутина 1 работает")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if isStoped {
				fmt.Println("Завершение работы канала 2")
				return
			} else {
				fmt.Println("Горутина 2 работает")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	stopChan <- 1 // отправляем сигнал о завершении первой горутины
	close(stopChan)
	isStoped = true // ставим флаг что вторая горутина тоже должна завершиться 

	wg.Wait()

}

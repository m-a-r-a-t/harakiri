/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataChannelEntity struct {
	ch       chan any
	isOpened bool
	mutex    sync.Mutex
}

type Worker struct {
	ch  <-chan any
	num int
	wg  *sync.WaitGroup
}

func main() {
	var timeCount int
	flag.IntVar(&timeCount, "t", 3, "Specify timeCount. Default is 3")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)
	dataChannelEntity := &DataChannelEntity{
		ch:       make(chan any),
		isOpened: true,
	}

	reader := &Worker{
		ch:  dataChannelEntity.ch,
		num: 1,
		wg:  &wg,
	}

	go reader.run()

	go runDataSender(dataChannelEntity, &wg)

	time.Sleep(time.Duration(timeCount) * time.Second) // ждем n секунд

	dataChannelEntity.mutex.Lock()
	dataChannelEntity.isOpened = false // закрываем канал для datasender
	close(dataChannelEntity.ch) // закрываем канал
	dataChannelEntity.mutex.Unlock()

	fmt.Println("Please wait until the data sender and data reader finish their work")

	wg.Wait() // ждем пока все горутины завершаться

}

func (w *Worker) run() {

	defer w.wg.Done()
	defer fmt.Println(fmt.Sprintf("Worker number %d closed", w.num))

	for data := range w.ch {
		fmt.Println("Worker num:", w.num, "Data:", data)
	}

	// эмуляция завершения работы
	time.Sleep(time.Second * 2)

}

func runDataSender(dataChannelEntity *DataChannelEntity, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Data sender closed")

	for {
		dataChannelEntity.mutex.Lock()
		if dataChannelEntity.isOpened {
			dataChannelEntity.ch <- rand.Int()

		} else {
			dataChannelEntity.mutex.Unlock()
			break
		}
		dataChannelEntity.mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
	}

	// эмуляция завершения работы
	time.Sleep(time.Second * 3)

}

/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct { 
	ch  <-chan any
	num int
	wg  *sync.WaitGroup
}

type DataChannelEntity struct { 
	ch       chan any
	isOpened bool // для того чтобы обозначить runDataSender ,что канал уже закрыт
	mutex    sync.Mutex // для безопасного доступа к DataChannelEntity
}

func main() {
	var countOfWorkers int

	signalChan, exit_chan := make(chan os.Signal, 1), make(chan int)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM) // подписываемся signalChan на события о завершении работы по нажатию Ctrl+C

	flag.IntVar(&countOfWorkers, "w", 1, "Specify countOfWorkers. Default is 1") // задаем число воркеров из stdinp

	flag.Parse()

	var workersWaitGroup sync.WaitGroup
	workersWaitGroup.Add(countOfWorkers + 1) // +1 так как помимо воркеров еще есть datasender
	dataChannelEntity := &DataChannelEntity{
		ch:       make(chan any),
		isOpened: true,
	}

	go interruptListener(signalChan, dataChannelEntity, &workersWaitGroup, exit_chan) 

	for i := 0; i < countOfWorkers; i++ {
		w := Worker{num: i + 1, ch: dataChannelEntity.ch, wg: &workersWaitGroup}
		go w.run() // запуск n воркеров 
	}

	go runDataSender(dataChannelEntity, &workersWaitGroup)

	exitCode := <-exit_chan // ждем сигнал о завершении
	os.Exit(exitCode)
}

func (w *Worker) run() {

	defer w.wg.Done()
	defer fmt.Println(fmt.Sprintf("Worker number %d closed", w.num))

	for data := range w.ch {
		fmt.Println("Worker num:", w.num, "Data:", data)
	}

	time.Sleep(time.Second * 2)

}

func runDataSender(dataChannelEntity *DataChannelEntity, wg *sync.WaitGroup) { // функция записи в канал
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

	time.Sleep(time.Second * 3)

}

func interruptListener(signalChan chan os.Signal, dataChannelEntity *DataChannelEntity, workersWaitGroup *sync.WaitGroup, exit_chan chan<- int) {
	for {
		s := <-signalChan

		if s == syscall.SIGINT {
			fmt.Println("Please wait until the workers finish their work")

			dataChannelEntity.mutex.Lock() // закрываем доступ к dataChannelEntity
			dataChannelEntity.isOpened = false // обозначаем что канал для datasender закрыт
			close(dataChannelEntity.ch) // закрываем канал: воркеры сами завершают свой цикл
			dataChannelEntity.mutex.Unlock() // открываем доступ к dataChannelEntity

			workersWaitGroup.Wait() // ждем пока все воркеры + datasender завершаться 
			break // выходим из цикла
		}

	}
	
	exit_chan <- 0 // отправляем сигнал о завершении
}

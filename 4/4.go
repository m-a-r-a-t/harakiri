// package main

// import (
// 	"flag"
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

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
	isOpened bool
	mutex    sync.Mutex
}

func main() {
	var countOfWorkers int

	signalChan, exit_chan := make(chan os.Signal, 1), make(chan int)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	flag.IntVar(&countOfWorkers, "w", 1, "Specify countOfWorkers. Default is 1")

	flag.Parse()

	var workersWaitGroup sync.WaitGroup
	workersWaitGroup.Add(countOfWorkers + 1)
	dataChannelEntity := &DataChannelEntity{
		ch:       make(chan any),
		isOpened: true,
	}

	go interruptListener(signalChan, dataChannelEntity, &workersWaitGroup, exit_chan)

	for i := 0; i < countOfWorkers; i++ {
		w := Worker{num: i + 1, ch: dataChannelEntity.ch, wg: &workersWaitGroup}
		go w.run()
	}

	go runDataSender(dataChannelEntity, &workersWaitGroup)

	exitCode := <-exit_chan
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

	time.Sleep(time.Second * 3)

}

func interruptListener(signalChan chan os.Signal, dataChannelEntity *DataChannelEntity, workersWaitGroup *sync.WaitGroup, exit_chan chan<- int) {
	for {
		s := <-signalChan

		if s == syscall.SIGINT {
			fmt.Println("Please wait until the workers finish their work")

			dataChannelEntity.mutex.Lock()
			dataChannelEntity.isOpened = false
			close(dataChannelEntity.ch)
			dataChannelEntity.mutex.Unlock()

			workersWaitGroup.Wait()
			break
		}

	}
	
	exit_chan <- 0
}

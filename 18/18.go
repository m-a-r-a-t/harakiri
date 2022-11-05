package main

import (
	"fmt"
	"sync"
)

type SyncCounter struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	var counter SyncCounter
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go increment(&counter, &wg)
	}
	wg.Wait()

	fmt.Println(counter.value)
}

func increment(counter *SyncCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 25; i++ {
		counter.mu.Lock()
		counter.value++
		counter.mu.Unlock()
	}

}

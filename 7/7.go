package main

import (
	"math/rand"
	"sync"
)

type SyncMap struct {
	data map[int]int
	mu   sync.Mutex
}

func main() {
	var wg sync.WaitGroup

	syncMap := &SyncMap{
		data: make(map[int]int),
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(syncMap, &wg)
	}

	wg.Wait()

}

func writeToMap(syncMap *SyncMap, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		syncMap.mu.Lock()
		syncMap.data[i] = rand.Int()
		syncMap.mu.Unlock()
	}
}

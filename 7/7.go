/*Реализовать конкурентную запись данных в map.*/

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

	stdSyncMap := sync.Map{}

	syncMap := &SyncMap{
		data: make(map[int]int),
	}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go writeToMap(syncMap, &wg) // запись в SyncMap.data
		go writeToStdSyncMap(&stdSyncMap, &wg) // запись в sync.Map
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

func writeToStdSyncMap(syncMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		syncMap.Store(i, rand.Int())
	}
}

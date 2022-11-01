package main

import (
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	syncMap := sync.Map{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(&syncMap, &wg)
	}

	wg.Wait()

}

func writeToMap(syncMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		syncMap.Store(i, rand.Int())
	}
}

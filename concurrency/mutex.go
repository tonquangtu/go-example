package main

import (
	"fmt"
	"sync"
)

var count = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	count = count + 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go increment(wg, mutex)
	}
	wg.Wait()
	fmt.Println("Result", count)
}

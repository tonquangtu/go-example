package main

import (
	"fmt"
	"sync"
	"time"
)

func call(wg *sync.WaitGroup, instance int)  {
	time.Sleep(1 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done()
}

func main() {
	fmt.Println("main started")
	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go call(&wg, i)
	}
	wg.Wait()
	fmt.Println("main stopped")
}

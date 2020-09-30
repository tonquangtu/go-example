package main

import (
	"fmt"
	"sync"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func worker(wg * sync.WaitGroup, jobs <-chan int, result chan<- int, instance int) {
	for num := range jobs {
		time.Sleep(10 * time.Millisecond)
		result <- num * num
		fmt.Printf("[worker-%v][job-%v] Sending result\n", instance, num)
	}
	wg.Done()
}

func main() {
	fmt.Println("[main] main started at", time.Since(start))
	wg := &sync.WaitGroup{}
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	jobCount := 6

	for i := 1; i<= 2; i++ {
		wg.Add(1)
		go worker(wg, jobs, results, i)
	}

	for i := 1; i<= jobCount; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	for i := 1; i <= jobCount; i++ {
		fmt.Println("Result", i, ":", <-results)
	}
	fmt.Println("[main] main() stopped at", time.Since(start))
}

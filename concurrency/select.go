package main

import (
	"fmt"
	"time"
)

func selectFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("for")
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("run")
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		time.Sleep(10^9)
		quit <- 0
	}()
	selectFibonacci(c, quit)
}

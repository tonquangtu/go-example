package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	fmt.Println("init")
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(30 * time.Millisecond)
	fmt.Println("service1() started at", time.Since(start))
	c <- "Result1"
}

func service2(c chan string) {
	time.Sleep(50 * time.Second)
	fmt.Println("service2() started at", time.Since(start))
	c <- "Result2"
}

func main() {
	fmt.Println("main() started at", time.Since(start))
	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	//time.Sleep(3 * time.Second)

	select {
	case val := <-chan1:
		fmt.Println("Response from service1", val, time.Since(start))
	case val := <-chan2:
		fmt.Println("Response from service2", val, time.Since(start))
	case <-time.After(30 * time.Millisecond):
		fmt.Println("No response received", time.Since(start))
		//default:
		//	fmt.Println("No response")
	}
	fmt.Println("main() stopped at", time.Since(start))
}

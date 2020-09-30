package main

import "fmt"

func main()  {
	fmt.Println("main started")
	c := make(chan int)
	go receiveOnly(c)
	c <- 10
	fmt.Println("main stopped")
}

func receiveOnly(c <-chan int) {
	fmt.Println("Receive from channel", <-c)
}

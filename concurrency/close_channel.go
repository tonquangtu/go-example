package main

import (
	"fmt"
)

func count(n int, c chan int) {
	for i := 0; i < n; i++ {
		fmt.Println("before goroutine")
		c <- i
		fmt.Println("after goroutine ")
	}
	close(c)
}

func main() {
	c := make(chan int)
	go count(5, c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	//time.Sleep(1000 * time.Millisecond)
}

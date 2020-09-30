package main

import (
	"fmt"
)

func greet(c chan string)  {
	fmt.Println("start greet")
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main started")
	c := make(chan string)
	//go greet(c)
	fmt.Println("before write to channel")
	c <- "Tu"
	fmt.Println("main stopped")
}

package main

import "fmt"

func main() {
	fmt.Println("main started")
	c := make(chan int, 3)
	//go squares(c)
	c <- 1
	c <- 2
	//c <- 3
	//c <- 4
	fmt.Printf("Length %v - Cap %v\n", len(c), cap(c))
	fmt.Println("main stopped")
}

func squares(c chan int) {
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num: ", num * num)
	}
}

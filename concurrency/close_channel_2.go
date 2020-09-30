package main

import (
	"fmt"
	"time"
)

func receive(c chan string)  {
	var val string
	var ok bool
	val, ok = <-c
	fmt.Println(val, ok)
	val, ok = <-c
	fmt.Println(val, ok)
	val, ok = <-c
	fmt.Println(val, ok)
}

func main() {
	fmt.Println("main started")

	c := make(chan string)

	go receive(c)

	c <- "Tu"
	//c <- "Ton"

	close(c)

	//val, ok := <-c
	//if val == "" {
	//	fmt.Println("nil")
	//	var s string
	//	if len(s) == 0 {
	//		fmt.Println("nil2")
	//	}
	//
	//}
	//fmt.Println(val, ok)

	//c <- "Close"

	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stopped")
}

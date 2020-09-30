package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go sender(c)

	for val := range c {
		fmt.Printf("Len %v - Val %v\n", len(c), val)
	}
}

func sender(c chan int) {
	c<-1
	c<-2
	c<-3
	c<-4
	close(c)
}

package main

import (
	"fmt"
	"time"
)

var start time.Time

func main() {
	fmt.Println("Main execution started at time", time.Since(start))
	//go printHello()
	//time.Sleep(10 * time.Millisecond)
	//fmt.Println("Main execution stopped")

	go printChars("HelloTu")
	go printDigits([] int32{1, 2, 3, 4})
	time.Sleep(200 * time.Millisecond)

	fmt.Println("\n main execution stopped at time", time.Since(start))
}

func init() {
	start = time.Now()
}

func printHello() {
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Hello World")
}

func printChars(s string) {
	for _, c := range s {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func printDigits(s []int32) {
	for _, n := range s {
		fmt.Printf("%d at time %v\n", n, time.Since(start))
		time.Sleep(30 * time.Millisecond)
	}
}

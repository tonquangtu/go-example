package main

import (
	"fmt"
	//"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := map[string]int{}
	fmt.Println(strings.Split(s, " "))
	for _, word := range strings.Split(s, " ") {
		counts[word] += 1
	}
	fmt.Println(counts)
	return counts
}

func main() {
	WordCount("I ate a donut. Then I ate another donut.")
	//wc.Test(WordCount)
}

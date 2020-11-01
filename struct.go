package main

import (
	"fmt"
)

type A struct {
	x int
}

type T struct {
	n int
	f float64
	A
	//time.Time
}

type TI interface {
	Get() int
	Inc(int)
}

type Container struct {
	TI
}

func (t *T) Get() int {
	return 1
}

func (t *T) Inc(i int) {
	t.n += i
}

func CreateTI() TI {
	t := &T{}
	defer func() {
		t.n = 3
	}()
	return t
}

func main() {
	var a A
	a.x = 1
	b := A{}
	fmt.Println(a)
	fmt.Println(b)
	if a == b {
		fmt.Println("Equals")
	}
	var t T
	t.n = 1
	t.A = a
	fmt.Println(t)

	var a2 = t.A
	a2.x = 2
	fmt.Println(a2, t)

	//structReturn := testStructReturn()
	//structReturn.n ++
	//fmt.Println(structReturn)
	//
	container := Container{CreateTI()}
	container.TI.Inc(4)
	ti2 := container.TI
	ti2.Inc(4)
	fmt.Println(container.TI)
}

func testStructReturn() *T {
	var t T
	defer func() {
		t.n = 2
	}()
	t.n = 1
	return &t
}

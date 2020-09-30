package main


import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Print()
	var x = 5
	fmt.Println(x)
	fmt.Println(add(2, 4))

	a, b := swap("Ton", "Tu")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var c, d, e bool
	e = true
	println(c, d, e)

	var v1, v2, v3 = 3, "tu", false
	println(v1, v2, v3)

	var v4 int
	var v5 string
	println(v4, v5)

	zeroValues()

	typeConversion()

	testBigNumber()

	testSlice()

}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}


func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func zeroValues()  {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q %v\n", i, f, b, s, len(s))
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// var i int = f, error need to explicit type conversion
	// var f2 float64 = x, error need to explicit type conversion
}

func testBigNumber()  {
	var bigInt big.Int
	bigInt.Exp(big.NewInt(2), big.NewInt(100), nil)
	fmt.Println(bigInt.String())

	const bigFloat  = 1<<100
	fmt.Printf("%f",float64(bigFloat))
	//fmt.Println(big.NewFloat(bigFloat).String())
}

func testSlice() {
	slice := []int {1, 2}
	slice1 := append(slice, 3)
	slice2 := append(slice1, 4)
	slice2[0] = -1
	printSlice(slice)
	printSlice(slice1)
	printSlice(slice2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


package benchmark

import (
	"testing"
)

//func TestFib(t *testing.T) {
//	fmt.Println(Fib(20))
//	fmt.Println(FibDP(20))
//}
//
//func BenchmarkFib(b *testing.B) {
//	fmt.Println(b.N)
//	for n := 0; n < b.N; n ++ {
//		Fib(20)
//	}
//}
//
//func BenchmarkFibDP(b *testing.B) {
//	fmt.Println(b.N)
//	for n := 0; n < b.N; n++ {
//		FibDP(20)
//	}
//}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

/**
https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
 */
func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }


package benchmark

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FibDP(n int) int {
	if n < 2 {
		return n
	}

 	f0, f1, f2 := 0, 1, 1
	for i := 2; i <= n; i++ {
		f2 = f1 + f0
		f0 = f1
		f1 = f2
	}
	return f2
}


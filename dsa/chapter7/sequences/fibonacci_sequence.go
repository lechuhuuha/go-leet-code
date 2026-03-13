package sequences

func Series(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}

	f[0] = 0
	f[1] = 1
	for i := 2; i < 2; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func FibonacciNumber(n int) int {
	if n <= 1 {
		return 1
	}

	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

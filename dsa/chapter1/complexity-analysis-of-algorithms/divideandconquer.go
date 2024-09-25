package main

// fibonacci method given k integer
func fibonacci(k int) int {
	if k <= 1 {
		return 1
	}
	return fibonacci(k-1) + fibonacci(k-2)
}

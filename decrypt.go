package main

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))

	if k == 0 {
		return result
	}

	start, end := 1, k

	if k < 0 {
		start, end = len(code)-Abs(k), len(code)-1
	}

	initialSum := 0
	for i := start; i <= end; i++ {
		initialSum += code[i]
	}

	for i := 0; i < len(code); i++ {
		result[i] = initialSum
		// Update sum by subtracting the element at start
		// and adding the element at end + 1,
		// using modulo to handle wrapping around the array.
		initialSum -= code[start%len(code)]
		initialSum += code[(end+1)%len(code)]

		start++
		end++
	}

	return result
}

// Helper function to calculate the absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

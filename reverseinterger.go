package main

import "math"

func ReverseInteger(x int) int {
	// Define the maximum and minimum 32-bit signed integer values
	maxInt32 := int(math.Pow(2, 31) - 1)
	minInt32 := int(math.Pow(-2, 31))
	result := 0
	for x != 0 {
		pop := x % 10
		x /= 10

		// Check for overflow before updating the result
		if result > maxInt32/10 || (result == maxInt32/10 && pop > 7) {
			return 0
		}
		if result < minInt32/10 || (result == minInt32/10 && pop < -8) {
			return 0
		}
		result = result*10 + pop
	}
	return result
}

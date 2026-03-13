package main

// Minimize the maximum number of items per store
func minimizedMaximum(n int, quantities []int) int {
	left, right := 1, 0

	// Find the maximum element in quantities to set the upper bound
	for _, quantity := range quantities {
		if quantity > right {
			right = quantity
		}
	}

	// Binary search to minimize the maximum number of items per store
	for left < right {
		middle := (left + right) / 2
		if canDistribute(quantities, middle, n) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}

func canDistribute(quantities []int, x, n int) bool {
	j := 0
	remaining := quantities[j]
	for i := 0; i < n; i++ {
		if remaining <= x {
			j++
			if j == len(quantities) {
				return true
			}
			remaining = quantities[j]
		} else {
			remaining -= x
		}
	}

	return false
}

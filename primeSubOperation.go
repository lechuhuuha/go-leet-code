package main

import "math"

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// primeSubOperation performs the operations as described.
func primeSubOperation(nums []int) bool {
	// Find the maximum element in the array.
	maxElement := math.MinInt
	for _, num := range nums {
		if num > maxElement {
			maxElement = num
		}
	}

	// Precompute the largest prime less than or equal to each number.
	previousPrime := make([]int, maxElement+1)
	for i := 2; i <= maxElement; i++ {
		if isPrime(i) {
			previousPrime[i] = i
		} else {
			previousPrime[i] = previousPrime[i-1]
		}
	}

	// Process each number in the array.
	for i := 0; i < len(nums); i++ {
		var bound int
		if i == 0 {
			bound = nums[i]
		} else {
			bound = nums[i] - nums[i-1]
		}

		// If the bound is less than or equal to 0, return false.
		if bound <= 0 {
			return false
		}

		// Find the largest prime less than the bound and subtract it.
		largestPrime := previousPrime[bound-1]
		nums[i] -= largestPrime
	}

	return true
}

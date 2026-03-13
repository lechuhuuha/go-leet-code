package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)

	// Create slices to store products to the left and right of each element
	left := make([]int, n)
	right := make([]int, n)

	// Initialize the left and right product arrays
	left[0] = 1
	right[n-1] = 1
	// Calculate products to the left of each element
	// [1 0 0 0 0] left
	// [-1 1 0 -3 3] nums
	// nums
	fmt.Println(nums)
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// Calculate products to the right of each element
	for i := n - 2; i >= 0; i-- {
		fmt.Println(right)

		right[i] = right[i+1] * nums[i+1]

	}

	// Calculate the final result by multiplying left and right products
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

func TestMethod() int {
	for i := 0; i < 10; i++ {
		return 0
	}
	return 0
}

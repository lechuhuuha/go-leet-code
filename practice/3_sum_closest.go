package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// Sort the array
	sort.Ints(nums)

	// Initialize variables
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			// Update closest sum
			if math.Abs(float64(target-currentSum)) < math.Abs(float64(target-closestSum)) {
				closestSum = currentSum
			}

			// Adjust pointers based on comparison with target
			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				// If the sum is equal to the target, it's the closest possible, so return it.
				return closestSum
			}
		}
	}

	return closestSum
}

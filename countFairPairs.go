package main

import (
	"slices"
)

// sort the slice in asc order

// for each number in the array, keep track of the smallest and largest numbers in the array
// that can form a fair pair with this number

// as you move to larger number, both boundaries move down
func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		high := lowerBound(nums, i+1, n-1, upper-nums[i]+1)
		low := lowerBound(nums, i+1, n-1, lower-nums[i])
		ans += 1 * (high - low)
	}

	return int64(ans)
}

func lowerBound(nums []int, low, high, element int) int {
	for low <= high {
		middle := low + (high-low)/2
		if nums[middle] >= element {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return low
}

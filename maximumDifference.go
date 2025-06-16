package main

// 0 <= i < j < n and nums[i] < nums[j]
func maximumDifference(nums []int) int {
	// invert loop
	// [7,1,5,4]
	// i = 7, j = 1
	// diff = 6
	// 7 > 1

	// i = 1, j = 5
	// diff = 4
	// 1 < 4

	// i = 5, j = 4
	// diff = 1
	// 5 > 4

	// maxDiff = 4

	// [9,4,3,2]
	// i = 9, j = 4
	// diff = 5
	// 9 > 4

	// i = 4, j = 3
	// diff = 1
	// 4 > 3

	// i = 3, j = 2
	// diff = 1
	// 3 > 2

	// maxDiff = -1

	// [1,5,2,10]
	// 1 - 5 = 4
	// 1 < 5
	// 5 - 2 = 3
	// 5 > 2
	// 2 - 10 = 8
	// 2 < 10
	// 1 - 10 = 9
	// 1 < 10
	// maxDiff = 9

	// store the lowest value and highest value
	// first time is lowest and higest is i and j
	// compare the result with contraint
	// O(n)
	minVal := nums[0]
	result := -1
	for j := 0; j < len(nums); j++ {
		if nums[j] > minVal {
			result = max((nums[j] - minVal), result)
		} else {
			minVal = nums[j]
		}
	}
	return result
}

package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	// sort the nums asc
	// create mutiple array with size 3
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)
	for i := 0; i < n; i += 3 {
		arr := nums[i : i+3]
		if arr[2]-arr[0] > k {
			return [][]int{}
		}
		result = append(result, arr)
	}

	return result
}

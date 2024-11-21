package main

import "runtime"

func removeDuplicates(nums []int) int {
	runtime.GC()
	uni := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[uni] = nums[i]
			uni++

		}
	}
	return uni
}

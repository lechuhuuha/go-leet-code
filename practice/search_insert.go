package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle
		}

		if nums[middle] < target {
			left = middle + 1
		}
	}

	return left
}

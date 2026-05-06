package interview

// You are given a sorted integer slice nums in ascending order and an integer target.
// Return the index of target if it exists. Otherwise, return -1.
// Input: nums = []int{-1, 0, 3, 5, 9, 12}, target = 9
// Output: 4

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		middle := l + (r-l)/2
		if target > nums[middle] {
			l = middle + 1
		} else if target < nums[middle] {
			r = middle - 1
		} else {
			return middle
		}
	}

	return -1

}

package problems

// Given a sorted array of integers nums and an integer target, return the starting and ending position of target.
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
func searchRange(nums []int, target int) []int {
	lowAns := -1
	highAns := -1

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		testNum := nums[mid]
		if testNum == target {
			lowAns = mid
			high = mid - 1 // discard the right, check the left. Since now high is the previous char to mid
		} else if testNum < target {
			low = mid + 1
		} else if testNum > target {
			high = mid - 1
		}
	}

	low = 0
	high = len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		testNum := nums[mid]
		if testNum == target {
			highAns = mid
			low = mid + 1 // discard the left, because the low is now the next char after mid
		} else if testNum < target {
			low = mid + 1
		} else if testNum > target {
			high = mid - 1
		}
	}
	return []int{lowAns, highAns}
}

package main

// Sliding Window with Deque
func resultsArray(nums []int, k int) []int {
	length := len(nums)
	if length < k {
		return []int{}
	}
	ans := make([]int, length-k+1)

	indexQueue := []int{}

	for i := 0; i < length; i++ {
		if len(indexQueue) > 0 && indexQueue[0] < i-k+1 {
			indexQueue = indexQueue[1:]
		}

		if len(indexQueue) > 0 && nums[i] != nums[i-1]+1 {
			indexQueue = nil // Clear the deque
		}

		indexQueue = append(indexQueue, i)

		if i >= k-1 {
			if len(indexQueue) == k {
				ans[i-k+1] = nums[indexQueue[len(indexQueue)-1]]
			} else {
				ans[i-k+1] = -1 // Not valid, return -1
			}
		}
	}

	return ans
}

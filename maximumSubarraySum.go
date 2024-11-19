package main

func maximumSubarraySum(nums []int, k int) int64 {
	ans := 0
	l := len(nums)
	currentSum := 0
	begin, end := 0, 0
	numToIndex := map[int]int{}

	for end < l {
		currNum := nums[end]
		lastOccurance, ok := numToIndex[currNum]
		if !ok {
			// lastOccurance stores the most recent index where currNum was found in the array.
			// If currNum hasn't been seen before (not in numToIndex),
			// there is no "last occurrence."
			// In such a case, we need a default value to indicate that
			// currNum is not part of the current sliding window.
			lastOccurance = -1
		}

		// If a duplicate is found, the window must start after the last occurrence of the duplicate
		// (lastOccurance + 1).
		// This ensures every element in the window is unique.
		for begin <= lastOccurance || end-begin+1 > k {
			// As we move begin to the right (to remove duplicates or reduce the window size to k),
			// the number at nums[begin] is no longer part of the current subarray.
			// Subtracting it ensures currentSum remains accurate.
			currentSum -= nums[begin]
			begin++
		}

		// This map helps us detect duplicates in the current window
		// and determine how far back they are (lastOccurance).
		// This ensures the sliding window logic can efficiently maintain the "no duplicates" constraint.
		numToIndex[currNum] = end

		// As we expand the sliding window by moving end to the right,
		// the new number (nums[end]) becomes part of the current subarray.
		// To calculate the sum of this new subarray, we must add nums[end] to currentSum.
		currentSum += nums[end]

		// Checking end-begin+1 == k ensures we only update ans when the current window
		// satisfies the size constraint.
		if end-begin+1 == k {
			ans = max(ans, currentSum)
		}

		end++
	}

	return int64(ans)
}

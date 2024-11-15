package main

func findLengthOfShortestSubarray(arr []int) int {
	right := len(arr) - 1
	left := 0

	// Move the right pointer leftward to find the end of a sorted (non-decreasing) suffix.
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}

	// Initialize ans as right to handle edge case like array alreay sorted then only get last or first element
	ans := right

	// if left pointer still in asc order
	for left < right && (left == 0 || arr[left-1] <= arr[left]) {
		// Move the right pointer further to find a valid range where arr[left:right] is sorted.
		// This loop shifts right until arr[left] <= arr[right], keeping it a non-decreasing sequence.
		for right < len(arr) && arr[left] > arr[right] {
			right++
		}

		// answer is the length of the removed subarray = right pointer - left pointer -1 (minux 1 because array start at 0 index)
		// Update the answer as the minimum length of a subarray that can be removed.
		ans = min(ans, right-left-1)
		left++
	}

	return ans
}

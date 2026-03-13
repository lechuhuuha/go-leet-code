package main

import "math"

// Pair struct to hold the cumulative sum and index
type Pair struct {
	Sum   int64
	Index int
}

func shortestSubarray(nums []int, k int) int {
	n := len(nums)

	// Stack-like slice to store cumulative sums and their indices
	cumulativeSumStack := []Pair{{0, -1}}

	runningCumulativeSum := int64(0)
	shortestSubarrayLength := math.MaxInt32

	for i := 0; i < n; i++ {
		// Update cumulative sum
		runningCumulativeSum += int64(nums[i])

		// Remove entries from stack that are larger than current cumulative sum
		for len(cumulativeSumStack) > 0 && runningCumulativeSum <= cumulativeSumStack[len(cumulativeSumStack)-1].Sum {
			cumulativeSumStack = cumulativeSumStack[:len(cumulativeSumStack)-1]
		}

		// Add current cumulative sum and index to stack
		cumulativeSumStack = append(cumulativeSumStack, Pair{runningCumulativeSum, i})

		// Find candidate index using binary search
		candidateIndex := findCandidateIndex(cumulativeSumStack, runningCumulativeSum-int64(k))

		// If a valid candidate is found, update the shortest subarray length
		if candidateIndex != -1 {
			shortestSubarrayLength = min(shortestSubarrayLength, i-cumulativeSumStack[candidateIndex].Index)
		}
	}

	// Return -1 if no valid subarray found
	if shortestSubarrayLength == math.MaxInt32 {
		return -1
	}
	return shortestSubarrayLength
}

// Binary search to find the largest index where cumulative sum is <= target
func findCandidateIndex(nums []Pair, target int64) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid].Sum <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

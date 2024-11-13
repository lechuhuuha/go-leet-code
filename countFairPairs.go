package main

import (
	"slices"
)

type FairPair struct {
	smallest int
	largest  int
}

// sort the slice in asc order

// for each number in the array, keep track of the smallest and largest numbers in the array
// that can form a fair pair with this number

// as you move to larger number, both boundaries move down
func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)
	n := len(nums)
	ans := 0

	fairPairs := FairPair{}
	fairPairs.smallest = nums[0]
	fairPairs.largest = nums[n-1]

	for i := 0; i < n; i++ {
		sum := nums[fairPairs.smallest] + nums[fairPairs.largest]
		if lower <= sum && sum <= upper {
			ans++
		}

	}

	return int64(ans)
}

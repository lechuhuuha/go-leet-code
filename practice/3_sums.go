package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		// use 2 pointer technique with current i as boundary
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// after we found answer then we skip this left and right position
				left++
				right--

				// move left pointer until we dont found any duplicates values
				// examples : -1, -1, 0
				// skip one -1
				// same goes for right pointer
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return ans
}

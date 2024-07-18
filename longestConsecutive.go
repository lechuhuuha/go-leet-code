package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	result := 0

	sort.Ints(nums)
	fmt.Println(nums)
	oldValue := nums[0]
	for i := 0; i < len(nums); i++ {
		if oldValue != nums[i] && nums[i]-oldValue == 1 {
			fmt.Println(nums[i], oldValue)
			result++
		}
		if i > 0 {
			oldValue = nums[i]
		}
	}

	return result
}

func longestConsecutive2(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		m[n] = struct{}{}
	}
	var longest int
	for n := range m {
		if _, ok := m[n-1]; !ok {
			length := 1
			for _, ok := m[n+length]; ok; _, ok = m[n+length] {
				length++
			}
			longest = max(longest, length)
		}
	}
	return longest
}

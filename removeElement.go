package main

import "fmt"

func removeElement(nums []int, val int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ans] = nums[i]
			fmt.Println(nums[ans], ans)
			ans++
		}
	}
	return ans
}

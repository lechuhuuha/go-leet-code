package main

const mod = 1000000007

// 0 <= i < j < nums.length
// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
func countNicePairs(nums []int) int {
	count := 0
	differenceFreq := make(map[int]int)
	// if nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
	// then nums[i] - rev(nums[j]) == nums[j] - rev(nums[i])
	// so nums[i] -  rev(nums[i]) == nums[j] - rev(nums[j])
	for _, num := range nums {
		reversed := reverse(num)
		difference := num - reversed
		count += differenceFreq[difference]
		count %= mod
		differenceFreq[difference]++
	}

	return count
}

func reverse(x int) int {
	reveserd := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		reveserd = reveserd*10 + pop
	}
	return reveserd
}

package main

func maximumSubarraySum(nums []int, k int) int64 {
	ans := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		subArray := []int{}

		subArray = append(subArray, nums[i])
		subArray = append(subArray, nums[i+1])
		subArray = append(subArray, nums[i+2])
		ans = max(ans, sumSlice(subArray))
	}

	return int64(ans)
}

func sumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

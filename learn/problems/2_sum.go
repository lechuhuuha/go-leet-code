package problems

func twoSum(nums []int, target int) []int {
	seenMap := map[int]int{}

	for key, x := range nums {
		complement := target - x

		if idx, ok := seenMap[complement]; ok {
			return []int{idx, key}
		}

		seenMap[x] = key
	}

	return []int{}
}

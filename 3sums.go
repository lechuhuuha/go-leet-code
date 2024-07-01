package main

func threeSum(nums []int) [][]int {
	result := [][]int{}
	seen := make(map[[3]int]bool)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			complement := -nums[i] - nums[j]
			if _, found := seen[[3]int{nums[i], nums[j], complement}]; found {
				// Triplet already exists
				continue
			}

			// Check if complement exists in the remaining part of the array
			for k := j + 1; k < len(nums); k++ {
				if nums[k] == complement {
					triplet := []int{nums[i], nums[j], nums[k]}
					result = append(result, triplet)
					seen[[3]int{nums[i], nums[j], nums[k]}] = true
					break
				}
			}
		}
	}

	return result
}

package interview

// You are given an integer slice nums.
// Move all zeroes to the end of the slice while keeping the relative order of the non-zero elements.
// You must do it in-place.
// Input:  []int{0, 1, 0, 3, 12}
// Output: []int{1, 3, 12, 0, 0}

func moveZeroes(nums []int) []int {
	w := 0

	for _, value := range nums {
		if value != 0 {
			nums[w] = value
			w++
		}
	}

	for w < len(nums) {
		nums[w] = 0
		w++
	}

	return nums
}

package main

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, num := range nums {
		if seen[num] == struct{}{} {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}

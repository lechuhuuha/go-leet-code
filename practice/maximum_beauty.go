package main

import "sort"

// items[i] = [pricei, beautyi]
func maximumBeauty(items [][]int, queries []int) []int {
	ans := make([]int, len(queries))

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0] // Sort by price in descending order
	})

	maxItem := items[0][1]

	for i := 0; i < len(items); i++ {
		maxItem = max(maxItem, items[i][1])
		items[i][1] = maxItem
	}

	for key, query := range queries {
		ans[key] = beautyBinarySearch(items, query)
	}

	return ans
}

func beautyBinarySearch(items [][]int, targetPrice int) int {
	left, right := 0, len(items)-1
	maxBeauty := 0

	for left <= right {
		mid := (left + right) / 2
		if targetPrice < items[mid][0] {
			right = mid - 1
		} else {
			maxBeauty = max(maxBeauty, items[mid][1])
			left = mid + 1
		}
	}
	return maxBeauty
}

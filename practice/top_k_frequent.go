package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)

	// Count frequencies of each element
	// Create a slice of unique elements
	// Sort unique elements based on their frequencies
	// Take the top K frequent elements

	for _, num := range nums {
		frequencyMap[num]++
	}

	uniqueElements := make([]int, 0, len(frequencyMap))
	for key := range frequencyMap {
		uniqueElements = append(uniqueElements, key)
	}

	sort.Slice(uniqueElements, func(i, j int) bool {
		return frequencyMap[uniqueElements[i]] > frequencyMap[uniqueElements[j]]
	})

	result := uniqueElements[:k]

	return result
}

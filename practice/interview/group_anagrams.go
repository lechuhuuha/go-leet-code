package interview

import "sort"

// Given a slice of strings, group the anagrams together.
// Input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// Output: [][]string{
//     {"eat", "tea", "ate"},
//     {"tan", "nat"},
//     {"bat"},
// }
// Order of groups does not matter. Order inside each group does not matter.

func groupAnagrams(strs []string) [][]string {
	sliceMap := map[string][]string{}
	for _, value := range strs {
		b := []byte(value)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		sliceMap[key] = append(sliceMap[key], value)
	}

	result := [][]string{}
	for _, slice := range sliceMap {
		result = append(result, slice)
	}
	return result
}

func groupAnagramsOptimized(strs []string) [][]string {
	sliceMap := map[[26]int][]string{}
	for _, value := range strs {
		key := countKey(value)
		sliceMap[key] = append(sliceMap[key], value)
	}

	result := [][]string{}
	for _, slice := range sliceMap {
		result = append(result, slice)
	}
	return result
}

func countKey(s string) [26]int {
	var count [26]int

	for _, ch := range s {
		count[ch-'a']++
	}

	return count
}

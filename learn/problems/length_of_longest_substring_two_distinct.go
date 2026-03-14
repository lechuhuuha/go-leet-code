package problems

// Given a string s, find the length of the longest substring that contains at most 2 distinct characters.

// eceba
func lengthOfLongestSubstringTwoDistinct(s string) int {
	left := 0
	seenMap := map[rune]int{}
	maxLen := 0

	runes := []rune(s)
	for right := 0; right < len(runes); right++ {
		char := runes[right]
		seenMap[char]++

		for len(seenMap) > 2 {
			leftChar := runes[left]
			seenMap[leftChar]--
			if seenMap[leftChar] == 0 {
				delete(seenMap, leftChar)
			}
			left++
		}

		length := right - left + 1
		if maxLen < length && len(seenMap) <= 2 {
			maxLen = length
		}
	}

	return maxLen
}

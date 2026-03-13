package problems

// pwwkew
func lengthOfLongestSubstring(s string) int {
	seenMap := map[rune]int{}
	maxLen := 0
	left := 0
	runes := []rune(s)
	for right := 0; right < len(runes); right++ {
		char := rune(s[right])
		if lastSeen, ok := seenMap[char]; ok && lastSeen >= left {
			left = seenMap[char] + 1
		}
		seenMap[char] = right
		length := right - left + 1
		if maxLen < length {
			maxLen = length
		}
	}
	return maxLen
}

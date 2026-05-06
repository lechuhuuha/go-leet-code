package interview

// Given a string s, return the length of the longest substring without repeating characters.
// Input:  "abcabcbb"
// Output: 3

// Explanation:
// The answer is "abc", length 3.

// Input:  "bbbbb"
// Output: 1

func lengthOfLongestSubstring(s string) int {
	left := 0
	result := 0
	seenMap := map[rune]int{}
	runes := []rune(s)
	for right := 0; right < len(runes); right++ {
		char := rune(runes[right])
		if lastSeen, ok := seenMap[char]; ok && lastSeen >= left {
			left = lastSeen + 1
		}

		seenMap[char] = right
		length := (right - left) + 1
		if length > result {
			result = length
		}
	}

	return result
}

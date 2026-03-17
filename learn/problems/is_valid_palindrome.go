package problems

// input: radar
// output: true
// input: abba
// output: true
func isStringPalindrome(s string) bool {
	if s == "" {
		return true
	}
	if len(s) == 1 {
		return true
	}
	left := 0
	runes := []rune(s)
	right := len(runes) - 1

	for left < right {
		if runes[left] == runes[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

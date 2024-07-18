package main

var digitToLetters = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	backtrack(digits, 0, "", &result)
	return result
}

func backtrack(digits string, index int, current string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, current)
		return
	}

	for _, letter := range digitToLetters[rune(digits[index])] {
		backtrack(digits, index+1, current+string(letter), result)
	}
}

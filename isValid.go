package main

func isValid(s string) bool {
	// the heap should alway be zero after for loop
	stack := []rune{}
	// Map to store closing to opening bracket mappings
	braceMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != braceMap[char] {
				return false
			}
			stack = stack[:len(stack)-1] // Pop from stack
		}
	}

	return len(stack) == 0
}

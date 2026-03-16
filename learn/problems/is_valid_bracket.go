package problems

type bracketStack []rune

func (s *bracketStack) Push(char rune) {
	*s = append(*s, char)
}

func (s *bracketStack) Pop() rune {
	if len(*s) == 0 {
		return 0
	}

	element := (*s)[len(*s)-1] // get the top or last element in the slice
	*s = (*s)[:len(*s)-1]      // create new slice that start at len of slice - 1 which mean we shrink the slice since we removed last element out of it
	return element
}

func (s bracketStack) Peek() rune {
	if len(s) == 0 {
		return 0
	}

	return s[len(s)-1]
}

func isValid(s string) bool {
	validStack := bracketStack{}
	openBracketMap := map[rune]struct{}{
		'(': {},
		'{': {},
		'[': {},
	}
	openCloseBracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		if _, ok := openBracketMap[char]; ok {
			validStack.Push(char)
		} else {
			peekOpenBracket := validStack.Peek()
			if openBrac, ok := openCloseBracketMap[char]; ok && openBrac == peekOpenBracket {
				validStack.Pop()
			} else {
				return false
			}
		}
	}

	return len(validStack) == 0
}

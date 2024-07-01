package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reveserd, original := 0, x
	for x != 0 {
		pop := x % 10
		x /= 10
		reveserd = reveserd*10 + pop
	}
	return reveserd == original
}

func isStringPalindrome(str string) bool {
	if len(str) < 1 {
		return false
	}
	var (
		runes []rune
	)
	str = strings.ToLower(str)
	for _, s := range str {
		if unicode.IsLetter(s) || unicode.IsNumber(s) { // Append the alphabet character to the result slice
			runes = append(runes, s)
		}
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(runes[j], runes[i])
		if runes[j] != runes[i] {
			return false
		}
	}

	return true
}

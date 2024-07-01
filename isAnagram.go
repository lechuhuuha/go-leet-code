package main

import (
	"sort"
	"strings"
)

func sortString(str string) string {
	// Convert the string to a slice of characters
	s := strings.Split(str, "")
	// Sort the slice of characters
	sort.Strings(s)
	// Join the sorted slice of characters back into a string
	return strings.Join(s, "")
}
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return sortString(s) == sortString(t)
}

func isAnagram(s string, t string) bool {
	// The rune type in Go represents a Unicode code point,
	// which is an integer value that uniquely identifies a character in the Unicode standard.
	// the trick is to store a map of unicode code point for the 2 string
	// if the map is the same then after add and minus 2 string then 2 string is the same
	if len(s) != len(t) {
		return false
	}
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		chars[v]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

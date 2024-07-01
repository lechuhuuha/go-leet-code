package main

import (
	"strings"
)

func Encode(ss []string) string {
	encodedString := strings.Join(ss, "4#")
	return encodedString
}

func Decode(s string) []string {
	decodedList := strings.Split(s, "4#")
	return decodedList
}

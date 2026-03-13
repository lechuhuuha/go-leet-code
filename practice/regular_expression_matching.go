package main

import "regexp"

func isMatch(s string, p string) bool {
	re, err := regexp.Compile("^" + p + "$")
	if err != nil {
		return false
	}
	return re.MatchString(s)
}

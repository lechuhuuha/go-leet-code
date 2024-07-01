// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"net/url"
)

func main() {

	// fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
	// s := "aa"
	// p := "a"
	// fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	// b := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }
	// fmt.Println(isValidSudoku(b))
	// fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	// str := []string{"Strings with spaces are tricky", "Another one with spaces", "This also has spaces"}

	fmt.Println(isStringPalindrome("0P"))
	// fmt.Println(isPalindrome(121))
	// fmt.Println(len(Decode(Encode(str))))

}

func GetURLField(d, field string) string {
	domain, err := url.Parse(d)
	result := ""
	if err != nil {
		return ""
	}
	switch field {
	case "scheme":
		result = domain.Scheme
	case "userinfo":
		result = domain.User.String()
	case "path":
		result = domain.Path
	case "query":
		result = domain.Query().Encode()
	case "fragment":
		result = domain.Fragment
	default:
		result = domain.Host
	}
	if result == "" {
		return domain.String()
	}
	return result
}

package spaceallocation

import "fmt"

func AddOne(num int) {
	num++
	fmt.Println("added to num", num, "Address of num", &num)
}

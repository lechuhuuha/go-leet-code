package main

import "fmt"

func h(x int, y int) int {
	return x * y
}

// g function which returns x and y parameters after modification
func g(l int, m int) (x int, y int) {
	x = 2 * l
	y = 4 * m
	return
}

// main method
func main() {
	fmt.Println(h(g(1, 2)))
}

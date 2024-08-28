package main

import "fmt"

func powerSeries(a int) (int, int) {

	return a * a, a * a * a
}

func Tuples() {
	var square, cube int
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, " cube", cube)
}

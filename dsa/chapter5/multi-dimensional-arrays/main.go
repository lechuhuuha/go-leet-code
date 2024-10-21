package main

import (
	"math/rand"
)

func threedimensional() [2][2][2]int {
	threedimensionalArr := [2][2][2]int{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				threedimensionalArr[i][j][k] = rand.Intn(3)
			}
		}
	}
	return threedimensionalArr
}

func main() {
	arr := threedimensional()
	tensorUnfold(arr, 0)
}

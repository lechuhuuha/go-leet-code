package main

import "fmt"

func tensorUnfold(threedimensionalArr [2][2][2]int, mode int) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(threedimensionalArr[mode][i][j])
		}
	}
}

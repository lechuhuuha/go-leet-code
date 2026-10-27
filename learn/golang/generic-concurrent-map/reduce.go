package main

import (
	"context"
	"fmt"
)

func Reduce[T any, R any](
	input <-chan T,
	initial R,
	fn func(R, T) R,
) R {
	for value := range input {
		initial = fn(initial, value)
	}
	return initial
}

func useReduce(ctx context.Context) {
	numbers := []int{1, 2, 3, 4, 5}
	input := FromSlice(ctx, numbers)

	squared := Map(ctx, input, 3, func(n int) int {
		return n * n
	})

	total := Reduce(squared, 0, func(acc, n int) int {
		return acc + n
	})

	fmt.Println(total) // 55
}

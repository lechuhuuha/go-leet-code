package main

import (
	"context"
	"fmt"
	"sync"
)

func Filter[T any](ctx context.Context, inputChan <-chan T, workers int, predicate func(T) bool) <-chan T {
	out := make(chan T)
	wg := sync.WaitGroup{}
	if workers <= 0 {
		workers = 1
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case input, ok := <-inputChan:
					if !ok {
						return
					}

					if predicate(input) {
						select {
						case <-ctx.Done():
							return
						case out <- input:
						}
					}

				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func useFilter(ctx context.Context) {
	products := []Product{
		{1, "Laptop", 1200},
		{2, "Mouse", 25},
		{3, "Keyboard", -80},
		{4, "Monitor", 350},
		{5, "Chip", -350}}

	input := FromSlice(ctx, products)

	filtered := Filter(ctx, input, 3, func(product Product) bool {
		return product.Price >= 100
	})

	named := Map(ctx, filtered, 2, func(product Product) string {
		return product.Name
	})

	totalPrice := Reduce(
		FromSlice(ctx, products),
		0.0,
		func(total float64, product Product) float64 {
			return total + product.Price
		},
	)
	fmt.Println(totalPrice) // 55

	for result := range named {
		fmt.Println(result)
	}
}

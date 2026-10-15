package main

import (
	"context"
	"fmt"
	"sync"
)

func FromSlice[T any](ctx context.Context, values []T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for _, value := range values {
			select {
			case <-ctx.Done():
				return

			case out <- value:
			}
		}
	}()

	return out
}

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

func Map[T any, R any](
	ctx context.Context,
	inputChan <-chan T,
	workers int,
	fn func(T) R,
) <-chan R {
	out := make(chan R)
	wg := sync.WaitGroup{}
	if workers <= 0 {
		workers = 1
	}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for { // if im not range over the input chan, im only get 3 value since we init 3 worker only
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d: Context cancelled. Exiting.\n", i)
					return
				case input, ok := <-inputChan:
					if !ok {
						return
					}
					result := fn(input)
					fmt.Printf("Worker %d: Receive the input %v.\n", i, input)

					select {
					case <-ctx.Done():
						return
					case out <- result:
					}
				}
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

type User struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	ctx := context.Background()
	products := []Product{
		{1, "Laptop", 1200},
		{2, "Mouse", 25},
		{3, "Keyboard", 80},
		{4, "Monitor", 350},
	}

	input := FromSlice(ctx, products)

	filtered := Filter(ctx, input, 3, func(product Product) bool {
		return product.Price >= 100
	})

	named := Map(ctx, filtered, 2, func(product Product) string {
		return product.Name
	})

	for result := range named {
		fmt.Println(result)
	}
}

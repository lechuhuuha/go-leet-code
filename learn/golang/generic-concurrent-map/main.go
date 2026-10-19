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

func MapErr[T any, R any](
	ctx context.Context,
	input <-chan T,
	workers int,
	fn func(T) (R, error),
	cancel context.CancelFunc,
) <-chan Result[R] {
	out := make(chan Result[R])
	wg := sync.WaitGroup{}
	if workers <= 0 {
		workers = 1
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d: Context cancelled. Exiting.\n", i)
					return
				case input, ok := <-input:
					resultErr := Result[R]{}
					if !ok {
						fmt.Println("channel close or input not correct", input) // when go close the channel input, we will receive zero value as default
						return
					}
					fmt.Printf("Worker %d: Receive the input %v.\n", i, input)

					result, err := fn(input)
					if err != nil {
						resultErr.Err = err
						select {
						case <-ctx.Done():
							return
						case out <- resultErr:
						}

						cancel() // cancellation waits until the error is received
						// If the consumer is slow or stops reading, Worker 1 may block sending the error and never reach cancel().
						fmt.Printf("Worker %d Cancelling all workers...\n", i)

						return
					}
					resultErr.Value = result
					select {
					case <-ctx.Done():
						return
					case out <- resultErr:
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

func Merge[T any](
	ctx context.Context,
	inputs ...<-chan T,
) <-chan T {
	out := make(chan T)
	wg := sync.WaitGroup{}

	for i, inputCHan := range inputs {
		wg.Add(1)
		go func(i int, iChan <-chan T) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d: Context cancelled. Exiting.\n", i)
					return
				case input, ok := <-iChan:
					if !ok {
						fmt.Println("channel close or input not correct", input) // when go close the channel input, we will receive zero value as default
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- input:
					}
				}
			}
		}(i, inputCHan)
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

type Result[T any] struct {
	Value T
	Err   error
}

func calculateDiscount(product Product) (float64, error) {
	if product.Price < 0 {
		return 0, fmt.Errorf("invalid price for product %s with price %f", product.Name, product.Price)
	}

	return product.Price * 0.9, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// useMapErr(ctx, cancel)
	useMerge(ctx)
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

func useMapErr(ctx context.Context, cancel context.CancelFunc) {
	products := []Product{
		{1, "Laptop", 1200},
		{2, "Mouse", 25},
		{3, "Keyboard", -80},
		{4, "Monitor", 350},
		{5, "Chip", -350},
	}

	input := FromSlice(ctx, products)

	// filtered := Filter(ctx, input, 3, func(product Product) bool {
	// 	return product.Price >= 100
	// })

	results := MapErr(
		ctx,
		input,
		3,
		calculateDiscount,
		cancel,
	)

	for result := range results {
		if result.Err != nil {
			fmt.Println("error:", result.Err)
			continue
		}

		fmt.Println("discounted price:", result.Value)
	}

	totalPrice := Reduce(
		FromSlice(ctx, products),
		0.0,
		func(total float64, product Product) float64 {
			return total + product.Price
		},
	)
	fmt.Println("totalPrice ", totalPrice)
}

func useMerge(ctx context.Context) {
	orders := FromSlice(ctx, []int{1, 2, 3})
	invoices := FromSlice(ctx, []int{4, 5, 6})

	merged := Merge(ctx, orders, invoices)

	for value := range merged {
		fmt.Println(value)
	}
}

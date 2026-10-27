package main

import (
	"context"
	"fmt"
	"sync"
)

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

func useMerge(ctx context.Context) {
	orders := FromSlice(ctx, []int{1, 2, 3})
	invoices := FromSlice(ctx, []int{4, 5, 6})

	merged := Merge(ctx, orders, invoices)

	for value := range merged {
		fmt.Println(value)
	}
}

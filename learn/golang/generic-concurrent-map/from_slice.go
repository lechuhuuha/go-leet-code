package main

import "context"

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

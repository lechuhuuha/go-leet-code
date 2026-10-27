package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RunLimited[T any](
	ctx context.Context,
	items []T,
	maxConcurrent int,
	fn func(context.Context, T) error,
) []error {
	if maxConcurrent <= 0 {
		maxConcurrent = 1
	}
	errorsSlice := make([]error, len(items))
	wg := sync.WaitGroup{}
	sem := NewSemaphore(maxConcurrent, ctx)
	for i, item := range items {
		if err := ctx.Err(); err != nil {
			for i := 0; i < len(items); i++ {
				errorsSlice[i] = ctx.Err()
			}
			break
		}

		err := sem.Acquire()
		if err != nil {
			for i := 0; i < len(items); i++ {
				errorsSlice[i] = ctx.Err()
			}
			break
		}
		wg.Add(1)
		go func(item T, i int) {
			defer func() {
				sem.Release()
				wg.Done()
			}()
			select {
			case <-ctx.Done():
				errorsSlice[i] = ctx.Err()
			default:
				fmt.Println(item)
				err := fn(ctx, item)
				if err != nil {
					errorsSlice[i] = err
				}
			}
		}(item, i)
	}

	wg.Wait()
	return errorsSlice
}

func printRunLimitedResults(orders []int, errs []error) {
	successCount := 0

	for i, err := range errs {
		if err != nil {
			fmt.Printf("Order %d failed: %v\n", orders[i], err)
			continue
		}

		successCount++
	}

	fmt.Printf("Successful: %d/%d\n", successCount, len(orders))
}

// Test 1: All 10 orders should complete.
func useRunLimitedSuccess(ctx context.Context) {
	fmt.Println("\n=== Test 1: All orders succeed ===")

	orders := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	start := time.Now()

	errs := RunLimited(ctx, orders, 3,
		func(ctx context.Context, orderID int) error {
			fmt.Printf("Processing order %d\n", orderID)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
			}

			fmt.Printf("Finished order %d\n", orderID)
			return nil
		},
	)

	printRunLimitedResults(orders, errs)
	fmt.Printf("Total time: %v\n", time.Since(start))
}

// Test 2: Order 1 fails, but the other 9 should succeed.
func useRunLimitedWithError(ctx context.Context) {
	fmt.Println("\n=== Test 2: Independent errors ===")

	orders := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	errs := RunLimited(ctx, orders, 3,
		func(ctx context.Context, orderID int) error {
			fmt.Printf("Processing order %d\n", orderID)

			if orderID == 1 {
				return fmt.Errorf("invalid payload for order %d", orderID)
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
			}

			fmt.Printf("Finished order %d\n", orderID)
			return nil
		},
	)

	printRunLimitedResults(orders, errs)
}

// Test 3: Cancel while three orders are being processed.
func useRunLimitedCancellation(ctx context.Context) {
	fmt.Println("\n=== Test 3: Context cancellation ===")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	orders := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	timer := time.AfterFunc(150*time.Millisecond, func() {
		fmt.Println(">>> Cancelling context")
		cancel()
	})
	defer timer.Stop()

	start := time.Now()

	errs := RunLimited(ctx, orders, 3,
		func(ctx context.Context, orderID int) error {
			fmt.Printf("Started order %d\n", orderID)

			select {
			case <-ctx.Done():
				// Simulate cleanup before the goroutine exits.
				fmt.Printf("Cleaning up order %d\n", orderID)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Stopped order %d\n", orderID)
				return ctx.Err()

			case <-time.After(600 * time.Millisecond):
				fmt.Printf("Finished order %d\n", orderID)
				return nil
			}
		},
	)

	printRunLimitedResults(orders, errs)
	fmt.Printf("RunLimited returned after %v\n", time.Since(start))
}

// Test 4: Only one order may run at any moment.
func useRunLimitedSequential(ctx context.Context) {
	fmt.Println("\n=== Test 4: Concurrency limit = 1 ===")

	orders := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mu sync.Mutex
	active := 0
	maxObserved := 0

	errs := RunLimited(ctx, orders, 1,
		func(ctx context.Context, orderID int) error {
			mu.Lock()
			active++
			if active > maxObserved {
				maxObserved = active
			}
			fmt.Printf("Started order %d (active: %d)\n", orderID, active)
			mu.Unlock()

			defer func() {
				mu.Lock()
				active--
				fmt.Printf("Finished order %d (active: %d)\n", orderID, active)
				mu.Unlock()
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
				return nil
			}
		},
	)

	printRunLimitedResults(orders, errs)
	fmt.Printf("Maximum observed concurrency: %d (expected 1)\n", maxObserved)
}

// Run all four tests.
func useRunLimited(ctx context.Context) {
	useRunLimitedSuccess(ctx)
	useRunLimitedWithError(ctx)
	useRunLimitedCancellation(ctx)
	useRunLimitedSequential(ctx)
}

func RunLimitedWithSemaphore[T any](
	ctx context.Context,
	items []T,
	sem *boundedSemaphore,
	fn func(context.Context, T) error,
) []error {
	errorsSlice := make([]error, len(items))
	wg := sync.WaitGroup{}

	for index, item := range items {
		if err := ctx.Err(); err != nil {
			// reason why we use index as the starting point
			// if error happen, we only write error to the index that
			// no go routine have ever claim
			// to prevent race condition map write
			for j := index; j < len(items); j++ {
				errorsSlice[j] = ctx.Err()
			}
			break
		}

		if err := sem.Acquire(ctx); err != nil {
			for j := index; j < len(items); j++ {
				errorsSlice[j] = ctx.Err()

			}
			break
		}

		if err := ctx.Err(); err != nil {
			sem.Release()
			for j := index; j < len(items); j++ {
				errorsSlice[j] = ctx.Err()
			}
			break
		}

		wg.Add(1)
		go func(i int, item T) {
			defer func() {
				sem.Release()
				wg.Done()
			}()
			select {
			case <-ctx.Done():
				errorsSlice[i] = ctx.Err()
			default:
				err := fn(ctx, item)
				if err != nil {
					errorsSlice[i] = err
				}
			}

		}(index, item)
	}

	wg.Wait()
	return errorsSlice
}

func useSharedSemaphore(ctx context.Context) {
	sem := NewUnboundedSemaphore(3)

	ctxA, cancelA := context.WithCancel(ctx)
	defer cancelA()

	ctxB, cancelB := context.WithCancel(ctx)
	defer cancelB()

	ordersA := []int{1, 2, 3, 4, 5, 6}
	ordersB := []int{101, 102, 103, 104, 105, 106}

	var mu sync.Mutex
	active := 0
	maxObserved := 0

	processOrder := func(ctx context.Context, orderID int) error {
		mu.Lock()
		active++
		if active > maxObserved {
			maxObserved = active
		}
		fmt.Printf("START order %d | active: %d\n", orderID, active)
		mu.Unlock()

		defer func() {
			mu.Lock()
			active--
			fmt.Printf("END   order %d | active: %d\n", orderID, active)
			mu.Unlock()
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(300 * time.Millisecond):
			return nil
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var errsA, errsB []error

	go func() {
		defer wg.Done()
		errsA = RunLimitedWithSemaphore(
			ctxA, ordersA, sem, processOrder,
		)
	}()

	go func() {
		defer wg.Done()
		errsB = RunLimitedWithSemaphore(
			ctxB, ordersB, sem, processOrder,
		)
	}()

	// Cancel only request A while both requests are processing.
	timer := time.AfterFunc(100*time.Millisecond, func() {
		fmt.Println(">>> Cancelling Request A")
		cancelA()
	})
	defer timer.Stop()

	wg.Wait()

	fmt.Printf("\nMaximum concurrency: %d (limit: 3)\n", maxObserved)

	fmt.Println("\nRequest A results:")
	printRunLimitedResults(ordersA, errsA)

	fmt.Println("\nRequest B results:")
	printRunLimitedResults(ordersB, errsB)
}

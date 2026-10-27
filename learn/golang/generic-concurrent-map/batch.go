package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Batch[T any](
	ctx context.Context,
	input <-chan T,
	batchSize int,
) <-chan []T {
	out := make(chan []T)
	wg := sync.WaitGroup{}
	wg.Add(1)
	if batchSize <= 0 {
		batchSize = 1
	}
	go func() {
		batchSlice := []T{}

		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Context cancelled. Exiting.\n")
				return
			case result, ok := <-input:
				if !ok {
					fmt.Println("channel close or input not correct", input) // when go close the channel input, we will receive zero value as default
					if len(batchSlice) > 0 {
						select {
						case <-ctx.Done():
							return
						case out <- batchSlice:
						}
					}
					return
				}

				batchSlice = append(batchSlice, result)
				// fmt.Println("result", result, batchSlice)
				if len(batchSlice) == batchSize {
					batchSliceCopy := make([]T, len(batchSlice))
					copy(batchSliceCopy, batchSlice)
					select {
					case <-ctx.Done():
						return
					case out <- batchSliceCopy:
						batchSlice = batchSlice[:0]

					}

				}

			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func BatchWithInterval[T any](
	ctx context.Context,
	input <-chan T,
	batchSize int,
	flushInterval time.Duration,
) <-chan []T {
	out := make(chan []T)
	if batchSize <= 0 {
		batchSize = 1
	}

	if flushInterval <= 0 {
		flushInterval = time.Second
	}

	go func() {
		batchSlice := make([]T, 0, batchSize)
		var timer *time.Timer
		var timerCh <-chan time.Time
		defer func() {
			close(out)
			if timer != nil {
				timer.Stop()
			}
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Context cancelled. Exiting.\n")
				return
			case result, ok := <-input:
				if !ok {
					fmt.Println("channel close or input not correct", input) // when go close the channel input, we will receive zero value as default
					if len(batchSlice) > 0 {                                 // final flush
						select {
						case <-ctx.Done():
							fmt.Printf("Context cancelled. Exiting.\n")
							return
						case out <- batchSlice:
						}
					}
					return
				}
				if len(batchSlice) == 0 {
					timer = time.NewTimer(flushInterval) // start the timer only when the first item arrives
					timerCh = timer.C
				}
				batchSlice = append(batchSlice, result)
				if len(batchSlice) == batchSize {
					batchSliceCopy := make([]T, len(batchSlice))
					copy(batchSliceCopy, batchSlice)
					select {
					case <-ctx.Done():
						fmt.Printf("Context cancelled. Exiting.\n")
						return
					case out <- batchSliceCopy:
						batchSlice = batchSlice[:0]
						timer.Stop()  // Stop the current timer when the batch is full.
						timerCh = nil // Disable the timer case until the next batch starts.
					}
				}

			case <-timerCh:
				fmt.Println("Timer fired, attempting to flush...")
				if len(batchSlice) > 0 {
					batchSliceCopy := make([]T, len(batchSlice))
					copy(batchSliceCopy, batchSlice)
					select {
					case <-ctx.Done():
						fmt.Printf("Context cancelled. Exiting.\n")
						return
					case out <- batchSliceCopy:
						batchSlice = batchSlice[:0]
						fmt.Println("flushed...")

					}
				}
				timerCh = nil
			}

		}
	}()
	return out
}

func useBatch(ctx context.Context) {
	input := FromSlice(ctx, []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	})

	batches := Batch(ctx, input, 3)

	for batch := range batches {
		fmt.Println(batch)
	}
}

func useBatchWithInterval(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	input := make(chan int)

	go func() {
		defer close(input)

		for _, n := range []int{1, 2, 3} {
			select {
			case <-ctx.Done():
				return
			case input <- n:
			}
		}

		// Keep the input open until cancellation.
		<-ctx.Done()
	}()

	start := time.Now()

	batches := BatchWithInterval(
		ctx,
		input,
		100,
		2*time.Second,
	)

	// Simulate a slow consumer.
	fmt.Println("Consumer busy for 5 seconds...")
	time.Sleep(5 * time.Second)

	for batch := range batches {
		fmt.Printf(
			"Received after %v: %v\n",
			time.Since(start).Round(time.Millisecond),
			batch,
		)

		cancel()
	}
}

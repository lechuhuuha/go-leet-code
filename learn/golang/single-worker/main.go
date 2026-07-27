package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var ErrDivisionByZero = errors.New("division by zero")

type Job struct {
	ID       int
	Dividend int
	Divisor  int
}

type Result struct {
	WorkerID int
	JobID    int
	Output   int
	Err      error
}

/*
The main goroutine sends five jobs into a channel.
One worker reads jobs from the channel.
The worker calculates the square of each number.
The worker sends the result into another channel.
Main prints all results.
The program exits without deadlock or goroutine leaks.
*/
func main() {
	workerCount := 3
	jobs := []Job{
		{ID: 1, Dividend: 10, Divisor: 2},
		{ID: 2, Dividend: 20, Divisor: 4},
		{ID: 3, Dividend: 15, Divisor: 0},
		{ID: 4, Dividend: 18, Divisor: 3},
		{ID: 5, Dividend: 100, Divisor: 0},
		{ID: 6, Dividend: 9, Divisor: 3},
	}
	jobChan := make(chan Job)
	resultChan := make(chan Result)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	go func() { // producer
		defer close(jobChan)

		for _, j := range jobs {
			select {
			case <-ctx.Done():
				return
			case jobChan <- j:
			}
		}
	}()

	for workerID := 1; workerID <= workerCount; workerID++ {
		wg.Add(1)
		go func(id int) { // consumer
			defer wg.Done()
			worker(ctx, id, jobChan, resultChan)
		}(workerID)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for r := range resultChan {
		if errors.Is(r.Err, ErrDivisionByZero) {
			fmt.Printf(
				"worker %d failed job %d: %v\n",
				r.WorkerID,
				r.JobID,
				r.Err,
			)
			continue
		}
		if r.JobID == 1 {
			fmt.Println("canceled since output == 3", r.WorkerID)
			cancel()
			continue
		}
		fmt.Printf("worker %d processed  job %d result: %d\n", r.WorkerID, r.JobID, r.Output)
	}
}

func worker(ctx context.Context, workerID int, jobs <-chan Job, results chan<- Result) {
	// TODO:
	// 1. Read jobs until the jobs channel is closed.
	// 2. Calculate Value * Value.
	// 3. Send a Result into results.
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			result := Result{
				JobID:    job.ID,
				WorkerID: workerID,
			}

			if job.Divisor == 0 {
				result.Err = ErrDivisionByZero
			} else {
				result.Output = job.Dividend / job.Divisor
			}
			// select {
			// case <-ctx.Done():
			// 	return
			// case results <- result:
			// }
			results <- result
		}
	}

}

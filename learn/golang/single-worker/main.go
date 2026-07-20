package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	WorkerID int
	JobID    int
	Output   int
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
		{ID: 1, Value: 2},
		{ID: 2, Value: 4},
		{ID: 3, Value: 6},
		{ID: 4, Value: 8},
		{ID: 5, Value: 10},
	}
	jobChan := make(chan Job)
	resultChan := make(chan Result)
	wg := sync.WaitGroup{}
	go func() {
		for _, j := range jobs {
			jobChan <- j
		}
		close(jobChan)
	}()

	for workerID := 1; workerID <= workerCount; workerID++ {
		wg.Add(1)
		go func(id int) {
			worker(id, jobChan, resultChan)
			defer wg.Done()
		}(workerID)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for r := range resultChan {
		fmt.Printf("worker %d job %d result: %d\n", r.WorkerID, r.JobID, r.Output)
	}
}

func worker(workerID int, jobs <-chan Job, results chan<- Result) {
	// TODO:
	// 1. Read jobs until the jobs channel is closed.
	// 2. Calculate Value * Value.
	// 3. Send a Result into results.
	for job := range jobs {
		result := Result{
			JobID:    job.ID,
			Output:   job.Value * job.Value,
			WorkerID: workerID,
		}
		results <- result
	}

}

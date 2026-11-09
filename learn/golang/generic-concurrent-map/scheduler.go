package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	RequestID  string
	CustomerID string
	OrderID    int
}

type Scheduler struct {
	mu      sync.Mutex
	pending map[string][]Job
	ready   ReadySlice

	notify chan struct{}
}

type ReadySlice []string

func (s *ReadySlice) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	customerID := (*s)[0]
	(*s)[0] = ""

	*s = (*s)[1:]
	return customerID
}

func (s *ReadySlice) Push(customerID string) {
	*s = append(*s, customerID)
}

func (s *ReadySlice) Len() int {
	return len(*s)
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		pending: map[string][]Job{},
		ready:   ReadySlice{},
		mu:      sync.Mutex{},
		notify:  make(chan struct{}),
	}
}

func (s *Scheduler) Submit(job Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Add the job to the correct CustomerID queue.
	customerID := job.CustomerID
	// the purpose of this excersise it to maintain fairness for each customer
	// meaning each customer should a chance to run request when their request come
	if len(s.pending[customerID]) == 0 {
		// that why we only append first request of each customer into the ready slice
		s.ready.Push(customerID)
		// s.ready = append(s.ready, customerID)
	}

	s.pending[customerID] = append(s.pending[customerID], job)

	// Wake waiting workers, then create a new notification channel.
	close(s.notify)
	// each go routine close different chan, avoid panic: close of closed channel
	// Creating a new channel allows future workers to wait for the next notification.
	s.notify = make(chan struct{})
}

func (s *Scheduler) Next() (Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Pop()
}

func (s *Scheduler) WaitNext(ctx context.Context) (Job, bool) {
	for {
		s.mu.Lock()
		// ctx is return and we can check for the err on why
		if ctx.Err() != nil {
			s.mu.Unlock()
			return Job{}, false
		}
		// if the queue is ready to pop
		if s.ready.Len() > 0 {
			job, ok := s.Pop()
			s.mu.Unlock()
			return job, ok
		}
		// acquire the notify chan
		notify := s.notify
		// so the main reason its because we wake go routine by use close chan
		// so if the worker listen on the new chan
		// they never knew about the close one
		// which lead to them miss the notify
		// and all go worker might miss which make the notification lost
		s.mu.Unlock()

		select {
		case <-ctx.Done():
			return Job{}, false
		case <-notify:
		}
	}
}

func (s *Scheduler) Pop() (Job, bool) {
	// Select the next job using round-robin.
	// Return false if no jobs are pending.
	if s.ready.Len() == 0 {
		return Job{}, false
	}
	// pop the customerID
	// customerID := s.ready[0]
	// s.ready = s.ready[1:]
	customerID := s.ready.Pop()

	// get the first request of the current customer ready
	jobs := s.pending[customerID]
	job := jobs[0]
	jobs[0] = Job{}
	if len(jobs) == 1 {
		// No more pending jobs for this customer.
		delete(s.pending, customerID)
	} else {
		// pop the current request
		s.pending[customerID] = jobs[1:]
		// move customer to end
		s.ready.Push(customerID)
		// s.ready = append(s.ready, customerID)
	}
	return job, true
}

func UseScheduler() {
	scheduler := NewScheduler()
	jobs := []Job{
		{
			RequestID:  "1",
			CustomerID: "1",
			OrderID:    1,
		},
		{
			RequestID:  "2",
			CustomerID: "1",
			OrderID:    2,
		},
		{
			RequestID:  "3",
			CustomerID: "2",
			OrderID:    3,
		},
		{
			RequestID:  "4",
			CustomerID: "2",
			OrderID:    4,
		},
		{
			RequestID:  "5",
			CustomerID: "3",
			OrderID:    5,
		},
	}

	for _, job := range jobs {
		scheduler.Submit(job)
	}

	for i := 0; i < len(jobs); i++ {
		fmt.Println(scheduler.Next())
	}
}

func UseAnotherScheduler(ctx context.Context) {
	scheduler := NewScheduler()

	scheduler.Submit(Job{RequestID: "1", CustomerID: "A", OrderID: 1})
	scheduler.Submit(Job{RequestID: "2", CustomerID: "A", OrderID: 2})
	scheduler.Submit(Job{RequestID: "3", CustomerID: "B", OrderID: 3})
	scheduler.Submit(Job{RequestID: "4", CustomerID: "B", OrderID: 4})

	// A gets the first turn.
	job, _ := scheduler.Next()
	fmt.Println("Next:", job.OrderID)

	// A new customer arrives while scheduling is in progress.
	scheduler.Submit(Job{RequestID: "5", CustomerID: "C", OrderID: 5})

	for {
		job, ok := scheduler.Next()
		if !ok {
			break
		}

		fmt.Println("Next:", job.OrderID)
	}
}

func useWaitNext(ctx context.Context) {
	scheduler := NewScheduler()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			fmt.Printf("Worker %d waiting...\n", workerID)

			job, ok := scheduler.WaitNext(ctx)
			if !ok {
				fmt.Printf("Worker %d stopped waiting\n", workerID)
				return
			}

			fmt.Printf("Worker %d received order %d\n", workerID, job.OrderID)
		}(i)
	}

	time.Sleep(200 * time.Millisecond)

	scheduler.Submit(Job{
		RequestID:  "R1",
		CustomerID: "A",
		OrderID:    1,
	})

	wg.Wait()
}

func useConcurrentScheduler(ctx context.Context) {
	scheduler := NewScheduler()

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(orderID int) {
			defer wg.Done()

			scheduler.Submit(Job{
				RequestID:  fmt.Sprintf("R%d", orderID),
				CustomerID: fmt.Sprintf("C%d", orderID%5),
				OrderID:    orderID,
			})
		}(i)
	}

	wg.Wait()

	count := 0
	for count < 100 {
		_, ok := scheduler.WaitNext(ctx)
		if !ok {
			break
		}
		count++
	}

	fmt.Printf("Processed %d/100 jobs\n", count)
}

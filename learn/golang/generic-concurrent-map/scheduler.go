package main

import (
	"context"
	"fmt"
	"slices"
	"sync"
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
	}
}

func (s *Scheduler) Submit(job Job) {
	s.mu.TryLock()
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
}

func (s *Scheduler) Next() (Job, bool) {
	s.mu.TryLock()
	defer s.mu.Unlock()

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

func (s *Scheduler) RoundRobin() []string {
	if len(s.pending) == 0 {
		return []string{}
	}

	maxRows := 0
	for _, request := range s.pending {
		if maxRows < len(request) {
			maxRows = len(request)
		}
	}

	jobOrders := make([]string, 0)

	for i := 0; i < maxRows; i++ {
		keys := make([]string, 0, len(s.pending))
		for k := range s.pending {
			keys = append(keys, k)
		}
		slices.Sort(keys)

		for _, customerID := range keys {
			jobs := s.pending[customerID]
			if i >= len(jobs) {
				continue
			}

			jobOrders = append(jobOrders, customerID)
		}
	}

	return jobOrders
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

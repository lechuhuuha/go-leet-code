package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Job struct {
	RequestID  string
	CustomerID string
	OrderID    int
}

type Scheduler struct {
	mu         sync.Mutex
	pending    map[string][]Job
	ready      ReadySlice
	processing map[int]Job // OrderID -> claimed job

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
		pending:    make(map[string][]Job),
		ready:      ReadySlice{},
		mu:         sync.Mutex{},
		notify:     make(chan struct{}),
		processing: make(map[int]Job),
	}
}

func (s *Scheduler) Submit(job Job) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.submit(job)
}

func (s *Scheduler) submit(job Job) {
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

func (s *Scheduler) WaitReady(ctx context.Context) bool {
	for {
		s.mu.Lock()
		if ctx.Err() != nil {
			s.mu.Unlock()
			return false
		}
		// if the queue is ready to pop
		if s.ready.Len() > 0 {
			s.mu.Unlock()
			return true
		}
		notify := s.notify
		s.mu.Unlock()

		select {
		case <-ctx.Done():
			return false
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
	s.processing[job.OrderID] = job
	return job, true
}

// ACK marks a claimed job as successfully processed.
func (s *Scheduler) Ack(orderID int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.processing[orderID]; !ok {
		return false
	}
	delete(s.processing, orderID)
	return true
}

// NACK requeues a claimed job for another attempt.
func (s *Scheduler) Nack(orderID int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	job, ok := s.processing[orderID]
	if !ok {
		return false
	}
	delete(s.processing, orderID)

	s.submit(job)

	return true
}

func RunScheduler(
	ctx context.Context,
	scheduler *Scheduler,
	sem *unboundedSemaphore,
	workers int,
	fn func(context.Context, Job) error,
) {
	wg := sync.WaitGroup{}

	if workers <= 0 {
		workers = 1
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
					isReady := scheduler.WaitReady(ctx)
					if !isReady {
						fmt.Println("ctx finished")
						return
					}

					if err := sem.Acquire(ctx); err != nil {
						fmt.Println("err sem.Acquire", err)
						return
					}

					if err := ctx.Err(); err != nil {
						fmt.Println("err ctx.Err()", err)
						sem.Release()
						return
					}

					job, ok := scheduler.Next()
					if !ok {
						fmt.Println("no ready job, continue")
						sem.Release()
						continue
					}

					// This ensures sem.Release() executes whenever fn returns, including during panic unwinding.
					// An unrecovered panic would still terminate the program.
					func() {
						defer sem.Release()

						if err := fn(ctx, job); err != nil {
							fmt.Printf("Order %d failed: %v\n", job.OrderID, err)
							scheduler.Nack(job.OrderID)
						}
						scheduler.Ack(job.OrderID)
					}()
				}
			}
		}(i)

	}
	wg.Wait()
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

func useRunScheduler(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	scheduler := NewScheduler()
	sem := NewUnboundedSemaphore(3)

	jobs := []Job{
		{CustomerID: "A", OrderID: 1},
		{CustomerID: "A", OrderID: 2},
		{CustomerID: "A", OrderID: 3},
		{CustomerID: "A", OrderID: 4},
		{CustomerID: "A", OrderID: 5},
		{CustomerID: "B", OrderID: 101},
		{CustomerID: "B", OrderID: 102},
		{CustomerID: "C", OrderID: 201},
	}

	for _, job := range jobs {
		scheduler.Submit(job)
	}

	var completed atomic.Int32
	var active atomic.Int32
	var maxObserved atomic.Int32

	RunScheduler(ctx, scheduler, sem, 5,
		func(ctx context.Context, job Job) error {
			current := active.Add(1)
			defer active.Add(-1)

			// Track maximum concurrency.
			for {
				old := maxObserved.Load()
				if current <= old ||
					maxObserved.CompareAndSwap(old, current) {
					break
				}
			}

			fmt.Printf(
				"START customer=%s order=%d active=%d\n",
				job.CustomerID,
				job.OrderID,
				current,
			)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
			}

			fmt.Printf("DONE order=%d\n", job.OrderID)

			if completed.Add(1) == int32(len(jobs)) {
				cancel()
			}

			return nil
		},
	)

	fmt.Printf(
		"Completed: %d/%d\nMax concurrency: %d\n",
		completed.Load(),
		len(jobs),
		maxObserved.Load(),
	)
}

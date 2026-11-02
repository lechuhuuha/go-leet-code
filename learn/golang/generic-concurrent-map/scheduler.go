package main

type Job struct {
	RequestID  string
	CustomerID string
	OrderID    int
}

type Scheduler struct {
	pending map[string][]Job
	ready   []string
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		pending: map[string][]Job{},
		ready:   []string{},
	}
}

func (s *Scheduler) Submit(job Job) {
	// Add the job to the correct RequestID queue.
	key := job.RequestID
	s.pending[key] = append(s.pending[key], job)
}

func (s *Scheduler) Next() (Job, bool) {
	// Select the next job using round-robin.
	// Return false if no jobs are pending.

}

func (s *Scheduler) RoundRobin() []string {
	if len(s.pending) == 0 {
		return []string{}
	}

	maxRows := 0
	totalJobs := 0
	for _, request := range s.pending {
		totalJobs += len(request)
		if maxRows < len(request) {
			maxRows = len(request)
		}
	}
	jobOrders := make([]string, totalJobs)

	for i := 0; i < maxRows; i++ { // maxRows = 5, i = 4
		for j := 0; j < len(s.pending); j++ { // j = 0, A
			jobs := s.pending["j"] // A1   A2   A3  A4 A5
			if i >= len(jobs) {
				continue
			}

			jobOrders = append(jobOrders, jobs) // [A1, B101, C201, A2, B102, C202, A3, C203, A4, A5]
		}
	}

	return jobOrders
}

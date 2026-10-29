package main

type RequestJobs[T any] struct {
	RequestID string
	Jobs      []T
}

// A1   A2   A3  A4 A5
// B101 B102
// C201 C202 C203

func RoundRobin[T any](requests []RequestJobs[T]) []T {
	// this func purpose is to organize which func should run first in the returned order
	// the order of each request should be the same, if a come first then a need to execute first
	// need to maintain a map for request ID
	// if
	if len(requests) == 0 {
		return []T{}
	}

	maxRows := 0
	totalJobs := 0
	for _, request := range requests {
		totalJobs += len(request.Jobs)
		if maxRows < len(request.Jobs) {
			maxRows = len(request.Jobs)
		}
	}
	jobOrders := make([]T, totalJobs)

	for i := 0; i < maxRows; i++ { // maxRows = 5, i = 4
		for j := 0; j < len(requests); j++ { // j = 0, A
			jobs := requests[j].Jobs // A1   A2   A3  A4 A5
			if i >= len(jobs) {
				continue
			}

			jobOrders = append(jobOrders, jobs[i]) // [A1, B101, C201, A2, B102, C202, A3, C203, A4, A5]
		}
	}

	return jobOrders
}

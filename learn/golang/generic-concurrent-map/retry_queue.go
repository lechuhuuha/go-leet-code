package main

import (
	"time"
)

type RetryItem struct {
	CustomerID string
	OrderID    int
	Attempt    int
	RetryAt    time.Time
}

type RetryQueue struct {
	items []RetryItem
}

func (q *RetryQueue) Len() int {
	return len(q.items)
}

func (q *RetryQueue) Less(i, j int) bool {
	// i = 10:00:02 j = 10:00:05
	return q.items[i].RetryAt.Before(q.items[j].RetryAt)
}

func (q *RetryQueue) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *RetryQueue) Push(x any) {
	q.items = append(q.items, x.(RetryItem))
}

func (q *RetryQueue) Pop() any {
	old := q.items
	n := len(old)
	item := old[n-1]
	old[n-1] = RetryItem{} // Avoid memory leaks if RetryItem holds pointers
	q.items = old[0 : n-1]
	return item
}

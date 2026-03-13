package main

type CircularQueue struct {
	size  int
	nodes []interface{}
	head  int
	last  int
}

func NewQueue(num int) *CircularQueue {
	circularQueue := CircularQueue{size: num + 1, head: 0, last: 0}
	circularQueue.nodes = make([]interface{}, circularQueue.size)
	return &circularQueue
}

func (c *CircularQueue) IsUnUsed() bool {
	return c.head == c.last
}

func (c *CircularQueue) IsComplete() bool {
	return c.head == (c.last+1)%c.size
}

func (c *CircularQueue) Add(element interface{}) {
	if c.IsComplete() {
		panic("Queue is completely utilized")
	}
	c.nodes[c.last] = element
	c.last = (c.last + 1) % c.size
}

func (c *CircularQueue) MoveOneStep() (element interface{}) {
	if c.IsUnUsed() {
		return nil
	}

	element = c.nodes[c.head]
	c.head = (c.head + 1) % c.size
	return
}

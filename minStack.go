package main

import "math"

type MinStack struct {
	vals []int
	min  int // Store the current minimum value
}

func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt64, // Initialize min to maximum possible integer
	}
}

func (this *MinStack) Push(val int) {
	if val <= this.min {
		// Push current min onto stack before updating min
		this.min = val
	}
	this.vals = append(this.vals, val)
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}
	this.vals = this.vals[0 : len(this.vals)-1]
}

func (this *MinStack) Top() int {
	if len(this.vals) == 0 {
		return 0
	}
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	m := this.vals[0]
	for i := 1; i < len(this.vals); i++ {
		m = min(m, this.vals[i])
	}
	return m
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

package main

import (
	"context"
	"fmt"
)

type semaphore struct {
	sem chan struct{}
	ctx context.Context
}

func NewSemaphore(nums int, ctx context.Context) *semaphore {
	sem := make(chan struct{}, nums)
	return &semaphore{
		sem: sem,
		ctx: ctx,
	}
}

func (s *semaphore) Acquire() error {
	select {
	case <-s.ctx.Done():
		fmt.Printf("Context cancelled. Exiting Acquire.\n")
		return s.ctx.Err()
	case s.sem <- struct{}{}:
		return nil
	}
}

func (s *semaphore) Release() {
	<-s.sem
}

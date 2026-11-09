package main

import (
	"context"
	"fmt"
)

type semaphore struct {
	sem chan struct{}
	ctx context.Context
}

// Semaphore with a shared context.
// All Acquire calls use the same context.
// Cancelling it signals all waiting Acquire calls to stop.
// Use when all callers share the same lifecycle.
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

type boundedSemaphore struct {
	sem chan struct{}
}

// Semaphore with per-call context.
// Each Acquire uses the context provided by its caller.
// Cancelling one caller's context does not affect other callers.
// Use when callers need independent cancellation.
func NewUnboundedSemaphore(nums int) *boundedSemaphore {
	sem := make(chan struct{}, nums)
	return &boundedSemaphore{
		sem: sem,
	}
}

func (s *boundedSemaphore) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Printf("Context cancelled. Exiting Acquire.\n")
		return ctx.Err()
	case s.sem <- struct{}{}:
		return nil
	}
}

func (s *boundedSemaphore) Release() {
	<-s.sem
}

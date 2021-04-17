package main

import (
	"context"
	"errors"
	"time"
)

var JobRejectError = errors.New("Pool is full")

type Fn func() (interface{}, error)

func NewWorkerPool(size int) *WorkerPool {
	return &WorkerPool{
		size: size,
		jobs: make(chan int, size),
	}
}

type WorkerPool struct {
	size int
	jobs chan int
}

func (w *WorkerPool) Submit(ctx context.Context, fn Fn, timeout time.Duration) (chan interface{}, error) {
	result := make(chan interface{})
	ctx, _ = context.WithTimeout(ctx, timeout)
	go w.run(ctx, fn, result)
	return nil, nil
}

func (w *WorkerPool) run(ctx context.Context, fn Fn, result chan<- interface{}) error {
	v, err := fn()
	result <- v
	return err
}

func (w *WorkerPool) Close() error {
	close(w.jobs)
	return nil
}

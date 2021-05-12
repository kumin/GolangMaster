package main

import (
	"context"
	"errors"
	"time"
)

var JobRejectError = errors.New("Pool is full")

type Job struct {
	ctx      context.Context
	Callable Runnable
	Timeout  time.Duration
	result   chan interface{}
}

func NewWorkerPool(size int) *WorkerPool {
	return &WorkerPool{
		size:      size,
		jobs:      make(chan Job, size),
		cancelFns: make([]context.CancelFunc, 0, size),
	}
}

type WorkerPool struct {
	size      int
	jobs      chan Job
	cancelFns []context.CancelFunc
}

func (w *WorkerPool) Submit(ctx context.Context, job Job) (chan interface{}, error) {
	if len(w.jobs) == w.size {
		return nil, JobRejectError
	}
	job.result = make(chan interface{})
	ctx, cancel := context.WithTimeout(ctx, job.Timeout)
	w.cancelFns = append(w.cancelFns, cancel)
	job.ctx = ctx
	w.jobs <- job
	return job.result, nil
}

func (w *WorkerPool) execute() {
	for j := range w.jobs {
		go w.run(j.ctx, j.Callable, j.result)
	}
}

func (w *WorkerPool) run(ctx context.Context, fn Runnable, result chan<- interface{}) error {
	v, err := fn.Run()
	result <- v
	return err
}

func (w *WorkerPool) Close() error {
	close(w.jobs)
	for _, fn := range w.cancelFns {
		fn()
	}
	return nil
}

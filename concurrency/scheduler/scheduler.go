package scheduler

import (
	"context"
	"log"
	"sync"
	"time"
)

func NewScheduler() *Scheduler {
	return &Scheduler{
		wg: new(sync.WaitGroup),
	}
}

type GracefulContext struct {
	ctx      context.Context
	cancelFn context.CancelFunc
}

type Scheduler struct {
	wg           *sync.WaitGroup
	jobs         []Runner
	gracefulCtxs []*GracefulContext
}

func (s *Scheduler) AddJob(ctx context.Context, job Runner) {
	s.jobs = append(s.jobs, job)
	ctx, cancel := context.WithCancel(ctx)
	s.gracefulCtxs = append(s.gracefulCtxs, &GracefulContext{ctx, cancel})
	s.wg.Add(1)
}

func (s *Scheduler) Process(ctx context.Context, job Runner) {
	defer s.wg.Done()
	ticker := time.NewTicker(job.IntervalTime())
	for {
		select {
		case <-ticker.C:
			job.Run(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	for idx, job := range s.jobs {
		go s.Process(s.gracefulCtxs[idx].ctx, job)
	}
}

func (s *Scheduler) GracefulShutdow(ctx context.Context) {
	log.Println("Sheduler is stopping...")
	for _, gracefulCtx := range s.gracefulCtxs {
		gracefulCtx.cancelFn()
	}
	s.wg.Wait()
	log.Println("Sheduler stopped")
}

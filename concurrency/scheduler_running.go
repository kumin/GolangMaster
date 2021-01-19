package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/my-packages/concurrency/scheduler"
)

var _ scheduler.Runner = &ExpCacheUpdater{}

func NewCacheUpdater(
	timeInterval time.Duration,
	jobId int,
) *ExpCacheUpdater {
	return &ExpCacheUpdater{
		TimeInterval: timeInterval,
		JobId:        jobId,
	}
}

type ExpCacheUpdater struct {
	TimeInterval time.Duration
	JobId        int
	IndexPath    string
}

func (e *ExpCacheUpdater) Run(ctx context.Context) error {
	fmt.Printf("Job %d is running after %f\n", e.JobId, e.TimeInterval.Minutes())
	return nil
}

func (e *ExpCacheUpdater) IntervalTime() time.Duration {
	return e.TimeInterval
}

func main() {
	ctx := context.Background()
	job1 := NewCacheUpdater(time.Duration(1)*time.Second, 1)
	job2 := NewCacheUpdater(time.Duration(1)*time.Second, 2)

	scheduler := scheduler.NewScheduler()
	scheduler.AddJob(ctx, job1)
	scheduler.AddJob(ctx, job2)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	scheduler.Start(ctx)

	<-quit
	scheduler.GracefulShutdow(ctx)
}

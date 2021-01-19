package scheduler

import (
	"context"
	"time"
)

type Runner interface {
	Run(ctx context.Context) error
	IntervalTime() time.Duration
}

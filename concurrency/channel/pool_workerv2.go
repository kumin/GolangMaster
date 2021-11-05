package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
)

type Runner interface {
	Run(ctx context.Context) error
}

type DoIt struct {
}

func (d *DoIt) Run(ctx context.Context) error {
	max := 0
	for i := 0; i <= 1000; i++ {
		randomNumber := rand.Intn(1000000000)
		if randomNumber > max {
			max = randomNumber
		}
	}
	fmt.Println(max)
	return nil
}

type Worker struct {
	ID          int
	Work        chan Runner
	PoolChannel chan chan Runner
	Done        chan bool
}

func (w *Worker) Execute(ctx context.Context) {
	log.Println("Start Worker: ", w.ID)
	go func() {
		for {
			w.PoolChannel <- w.Work
			select {
			case job := <-w.Work:
				job.Run(ctx)
			case <-w.Done:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.Done <- true
}

type Executor struct {
	PoolSize    int
	PoolChannel chan chan Runner
	Work        chan Runner
	Done        chan bool
}

func (e *Executor) Start(ctx context.Context) {
	log.Println("executor started with size: ", e.PoolSize)
	workers := make([]*Worker, 0, e.PoolSize)
	for i := 0; i < e.PoolSize; i++ {
		worker := &Worker{
			ID:          i,
			Work:        make(chan Runner),
			PoolChannel: e.PoolChannel,
			Done:        make(chan bool),
		}
		worker.Execute(ctx)
		workers = append(workers, worker)
	}

	go func() {
		for {
			select {
			case <-e.Done:
				for _, w := range workers {
					w.Stop()
				}
				return
			case job := <-e.Work:
				worker := <-e.PoolChannel
				worker <- job
			}
		}
	}()
}

func main() {
	executor := &Executor{
		PoolSize:    5,
		PoolChannel: make(chan chan Runner),
		Work:        make(chan Runner),
		Done:        make(chan bool),
	}
	executor.Start(context.Background())
	for i := 0; i <= 100; i++ {
		executor.Work <- &DoIt{}
	}
}

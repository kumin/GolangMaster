package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	apps "github.com/kumin/GolangMaster/restful/apps/server-ctl"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSYS)
	defer done()
	server, err := apps.BuildServer()
	if err != nil {
		log.Fatal(err)
	}
	_, err = apps.BuildMetricServer()
	if err != nil {
		log.Fatal(err)
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return server.Start(ctx)
	})
	if eg.Wait() != nil {
		panic(err)
	}
}

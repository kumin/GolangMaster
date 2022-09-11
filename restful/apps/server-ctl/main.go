package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"github.com/kumin/GolangMaster/restful/handler"
	"github.com/kumin/GolangMaster/restful/services"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSYS)
	defer done()
	server := NewHttpServer(8080)
	prodService := services.NewProductCtlHandler()
	prodHandler := handler.NewProductCtlHandler(prodService)
	server.RegisterHandler("/v1/product/listing", handler.HandlerWrapper(prodHandler.ListProducts))
	server.RegisterHandler("/v1/product", handler.HandlerWrapper(prodHandler.GetProduct))
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		go server.Start(ctx)
		select {
		case <-ctx.Done():
			cancel()
		}
		defer func() {
			wg.Done()
		}()
	}(ctx)
	wg.Wait()
}

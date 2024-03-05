package main

import (
	"sync"

	ramq "github.com/kumin/GolangMaster/msgqueue/rabbitmq"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ramq.PublishMsg()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ramq.ConsumeMsg()
	}()

	wg.Wait()
}

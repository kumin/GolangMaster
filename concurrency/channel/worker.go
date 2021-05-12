package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func worker(c <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-c)
	wg.Done()
}

type Console struct {
	content string
}

func (c *Console) Run() (interface{}, error) {
	fmt.Printf("exec task: %s", c.content)

	return c.content, nil
}

func ConsoleHandler(content interface{}) (interface{}, error) {
	return "", nil
}

func main() {
	ctx := context.Background()
	chanMap := make(map[int]chan interface{})
	workerPool := NewWorkerPool(9)
	for i := 0; i <= 5; i++ {
		ch, err := workerPool.Submit(ctx, Job{
			Callable: &Console{
				content: strconv.Itoa(i),
			},
			Timeout: 1 * time.Second,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		chanMap[i] = ch
	}
	for v, k := range chanMap {
		fmt.Printf("Result Task %d: %s", v, <-k)
	}
}

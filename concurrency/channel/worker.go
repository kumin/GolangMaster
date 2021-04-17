package main

import (
	"fmt"
	"sync"
)

func worker(c <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-c)
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	c := make(chan string, 1)
	//go worker(c, wg)
	wg.Add(1)
	//go worker(c, wg)
	//wg.Add(1)
	c <- "kumin"
	//c <- "haha"
	wg.Wait()
}

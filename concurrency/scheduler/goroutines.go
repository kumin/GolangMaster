package scheduler

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func TestRoutine(t *testing.T) {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Printf("counter = %d", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		fmt.Printf("Thread %d is increasing\n", id)
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}

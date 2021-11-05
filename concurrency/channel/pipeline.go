package main

import (
	"fmt"
	"testing"
	"time"
)

func Test() {
	generator := func(done <-chan interface{}, nums ...int) <-chan int {
		inStream := make(chan int)
		go func() {
			defer close(inStream)
			for _, i := range nums {
				select {
				case <-done:
					return
				case inStream <- i:
				}
			}
		}()
		return inStream
	}

	generator(nil, 1, 2, 3)
}

func Pipeline_Test(t *testing.T) {
	nums := make(chan int)
	go func() {
		defer close(nums)
		time.Sleep(3 * time.Second)
		for _, nu := range []int{0, 2, 3, 4} {
			nums <- nu
			fmt.Println("hoho")
		}
	}()
	for nu := range nums {
		fmt.Println(nu)
	}
}

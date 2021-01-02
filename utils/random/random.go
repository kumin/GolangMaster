package main

import (
	"fmt"
	"math/rand"
)

type RandInRFn func(limit int64) int64

func NewLimitRandom() RandInRFn {
	return func(limit int64) int64 {
		return rand.Int63n(limit)
	}
}

func main() {
	randFn := NewLimitRandom()
	fmt.Println(randFn(100))
}

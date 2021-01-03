package main

import (
	"fmt"
	"github.com/my-packages/helloworld/mypackage_test"
)

func main() {
	fmt.Println("I am a master Golang")
	fmt.Println(mypackage_test.Add(1, 2))
}

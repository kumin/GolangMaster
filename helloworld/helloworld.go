package main

import (
	"fmt"
	"github.com/my-packages/helloworld/mypackage"
)

func main() {
	fmt.Println("I am a master Golang")
	fmt.Println(mypackage.Add(1, 2))
}

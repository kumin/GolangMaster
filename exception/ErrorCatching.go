package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func Perform() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	GoesWrong()
	return
}

func GoesWrong() {
	panic(errors.New("Fail"))
}

func main() {
	err := Perform()
	fmt.Println(err)
}

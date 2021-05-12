package main

type Runnable interface {
	Run() (interface{}, error)
}

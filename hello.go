package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//by default,send and receiver blocks when the other side is not ready
	//so you can not write to and read from a same channel in same goroutine step by step, it causes deadlock.
	go write(c)
	fmt.Print(read(c))
}

func write(c chan int) {
	fmt.Println("write to channel")
	c <- 1
}

func read(c chan int) int {
	fmt.Println("read from channel")
	return <-c
}

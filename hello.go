package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	//by default,send and receiver blocks when the other side is not ready
	//so you can not write to and read from a same channel in same goroutine step by step, it causes deadlock.
	//and code below demostrate that read will not return until write func write a value to channel.
	go read(c)
	time.Sleep(2 * time.Second)
	go write(c)
	time.Sleep(10 * time.Millisecond)
}

func write(c chan int) {
	fmt.Println("write to channel")
	c <- 1
}

func read(c chan int) int {
	fmt.Println("come in read from channel")
	defer fmt.Println("read a value")
	return <-c
}

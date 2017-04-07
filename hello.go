package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//by default,send and receiver blocks when the other side is not ready
	//so you can not write to and read from a same channel in same goroutine step by step, it causes deadlock.
	//One can not guarantee a go func to be execute, maybe the main thread is over before CPU has time to call go func
	c := 12
	go write(c)
	go read(c)
	time.Sleep(10000 * time.Nanosecond)
}

func write(c int) {
	fmt.Println("write to channel" + strconv.Itoa(c))
}

func read(c int) int {
	fmt.Println("read from channel" + strconv.Itoa(c))
	return c
}

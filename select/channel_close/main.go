package main

import (
	"fmt"
	"time"
)

func main() {
	TestSelect1()
}

func TestSelect1() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(2 * time.Second)
		close(c)
		// c <- "hello"
	}()

	fmt.Println("Blocking on read...")
	select {
	case v := <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
		fmt.Println(v)
	}
}

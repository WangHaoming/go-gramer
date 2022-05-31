package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	// 等待时间相同的话，select会随机选择一个准备好的通道执行。
	go func() {
		time.Sleep(1000000 * time.Second)
		c1 <- 1
	}()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	c2 <- 2
	// }()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(<-c3)
	// }()
	select {
	case v1 := <-c1:
		fmt.Printf("received %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("received %v from c2\n", v2)
	case c3 <- 23:
		fmt.Printf("sent %v to c3\n", 23)
		// default:
		// 	fmt.Printf("no one was ready to communicate\n")
	}

}

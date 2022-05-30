package main

import (
	"fmt"
)

var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	}
)

func main() {
	// a = 1
	// f := "abc"
	// fmt.Println(f)

	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
	// fmt.Println(str)
}

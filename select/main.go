package main

import (
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go Getdata("https://www.google.jp", ch1)
	go Getdata("https://www.facebook.com", ch2)
	go Getdata("https://www.yahoo.jp", ch3)
	select {
	case v := <-ch1:
		fmt.Println(v)
	case <-ch2:
		fmt.Println("face")
	case <-ch3:
		fmt.Println("baidu")
	}
}

func Getdata(url string, ch chan string) {
	_, err := HttpRequest.Get(url)
	if err != nil {

	} else {
		ch <- url
	}
}

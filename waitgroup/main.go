package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func main() {
	tasks := []func(){
		func() {

			fmt.Printf("1. ")
		},
		func() {
			time.Sleep(time.Second)
			fmt.Printf("2. ")
		},
	}

	var wg sync.WaitGroup
	for _, task := range tasks {
		// Go 对 array/slice 进行遍历时，runtime 会把 task[i] 拷贝到 task 的内存地址，下标 i 会变，而 task 的内存地址不会变。如果不进行这次赋值操作，所有 goroutine 可能读到的都是最后一个 task。
		task := task
		wg.Add(1)
		go func() {
			task()
			fmt.Println()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
	for idx, task := range tasks {
		task()
		fmt.Printf("遍历 = %v, ", unsafe.Pointer(&task))
		fmt.Printf("下标 = %v, ", unsafe.Pointer(&tasks[idx]))
		task := task
		fmt.Printf("局部变量 = %vn", unsafe.Pointer(&task))
	}
}

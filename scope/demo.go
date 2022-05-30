// package main // 声明 main 包

// import "fmt"

// var d int // 全局变量
// func main() {
// 	var a int = 3
// 	var b int = 4
// 	c := a + b
// 	d = a + b
// 	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
// }

// 形参相当于函数的局部变量
package main

import (
	"fmt"
)

//全局变量 a
var a int = 13

func main() {
	//局部变量 a 和 b
	var a int = 3
	var b int = 4
	fmt.Printf("main() 函数中 a = %d\n", a)
	fmt.Printf("main() 函数中 b = %d\n", b)
	c := sum(a, b)
	fmt.Printf("main() 函数中 c = %d\n", c)
}
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

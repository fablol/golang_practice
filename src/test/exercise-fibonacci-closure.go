package main

import (
	"fmt"
)

// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		temp := a
		a, b = b, (a + b)
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

package main

import (
	"fmt"
	"runtime/debug"
)

var x, y = 0, 9

func ThorwsError(x, y int) {
	defer func() {
		ret := recover()
		if ret != nil {
			fmt.Println("Code paniced:", ret)
			debug.PrintStack()
		}
	}()

	if x < 0 || y < 0 {
		panic("x or y is zero")
	}
	
	fmt.Println(x, y)
}

func main() {
	
	fmt.Println("Demo: Panic")

	ThorwsError(100, 10)
	ThorwsError(200, -1)
	ThorwsError(50, 10)
}

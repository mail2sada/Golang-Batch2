package main

import (
	"fmt"
	"runtime/debug"
)

func Divide(a, b int) (res int) {
	// defer func() {
	// 	rec := recover()
	// 	if rec != nil {
	// 		fmt.Println("Received error..", rec)
	// 	}
	// }()

	res = a / b
	return
}

func main() {
	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println("Received error..", rec)
			debug.PrintStack()
		}
	}()
	fmt.Println("Demo: Panic and recover")

	a := Divide(10, 20)
	fmt.Println("Result is", a)
	b := Divide(50, 0)
	fmt.Println("Result is ###", b)
}

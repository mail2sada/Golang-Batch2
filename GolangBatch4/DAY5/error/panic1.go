package main

import (
	"fmt"
)

func Divide(a, b int) int {
	if b == 0 {
		panic("divide by zero....")
	}

	return a / b
}

func main() {

	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println("Panic while executing main")
			//debug.PrintStack()
		}
	}()
	fmt.Println("Panic")

	Divide(10, 2)

	Divide(10, 0)
}

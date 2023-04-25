package main

import (
	"fmt"
	"reflect"
)

func Add(a int, b int) int {
	sum := a + b
	return sum
}

func Sum(a, b float64) int {
	sum := a + b
	return int(sum)
}

func main() {
	fmt.Println("Demo: Function return value")

	ans := Add(10, 20)

	fmt.Println("Result is ", ans)

	fmt.Println("Type of ans is", reflect.TypeOf(ans))

	fmt.Println("Sum of 40, 50", Sum(40.4, 50.0))
}

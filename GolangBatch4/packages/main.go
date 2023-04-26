package main

import (
	"fmt"
	"fun/math"
)

func main() {
	a := math.Add(10, 20)
	fmt.Println("Value of a is ", a)

	a = math.Multiply(10, 5)
	fmt.Println("Value of a is ", a)
	
}

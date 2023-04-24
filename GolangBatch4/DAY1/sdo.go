package main

import "fmt"

func main() {
	fmt.Println("Demo: Short declaration operator")
	var a = 100

	b := 200
	fmt.Println("Value of a b", a, b)

	x, y := 10.33, "Hello"

	fmt.Println(x, y)
}

package main

import "fmt"

func Add(a, b int) (res int) {
	res = a + b
	return
}

func main() {
	fmt.Println("Demo: Functions named return")

	fmt.Println("Value of 10, 20", Add(10, 20))
}

package main

import "fmt"

func Swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 10, 20

	a, b = Swap(a, b)

	fmt.Println(a, b)
}

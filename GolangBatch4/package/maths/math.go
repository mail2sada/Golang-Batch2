package math

import "fmt"

func init() {
	fmt.Println("1 init from math package")
}

func init() {
	fmt.Println("2 init from math package")
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) float64 {
	return float64(a) / float64(b)
}

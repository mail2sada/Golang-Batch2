package main

import "fmt"

type Details struct {
	int
	string
	float32
	float64
}

func main() {
	fmt.Println("Demo: Prompted structure")
	var detail = Details{int: 1, string: "Hello", float32: 3.142, float64: 45.445}
	fmt.Println(detail)
}

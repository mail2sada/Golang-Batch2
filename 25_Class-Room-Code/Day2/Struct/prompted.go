package main

import "fmt"

type prompted_struct struct {
	int
	string
	float32
	uint
	uint64
}

func main() {
	fmt.Println("Demo: Prompted fields")

	p := prompted_struct{int: 10, string: "hello", float32: 56.78, uint: 100, uint64: 200}

	fmt.Println(p)

	fmt.Println(p.int)
}

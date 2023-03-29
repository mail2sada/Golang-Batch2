package main

import (
	"fmt"
	"reflect"
)

type prompted_struct struct {
	int
	string
	float32
	uint
	uint64
}

type test struct {
	int
	string
}

func main() {
	fmt.Println("Demo: Prompted fields")

	p := prompted_struct{int: 10, string: "hello", float32: 56.78, uint: 100, uint64: 200}

	q := prompted_struct{int: 10, string: "hello", float32: 56.78, uint: 100, uint64: 200}

	z := test{int: 1, string: "test"}
	if p == q {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not equals..")
	}

	if reflect.DeepEqual(p, ) {
		fmt.Println("Equlas: DeepEqual")
	} else {
		fmt.Println("Not equal: DeepEqual")
	}
	fmt.Println(p)

	fmt.Println(p.int)
}

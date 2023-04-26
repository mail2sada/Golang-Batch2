package main

import "fmt"

type Test struct {
	int
}

func main() {

	var a interface{}

	a = 10

	fmt.Println(a)

	a = "Hello"

	fmt.Println(a)

	a = Test{}

	fmt.Println(a)
}

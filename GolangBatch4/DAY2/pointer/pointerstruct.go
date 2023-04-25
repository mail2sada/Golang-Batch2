package main

import "fmt"

type Test struct {
	int
	string
}

func main() {

	fmt.Println("Demo: Struct pointers")
	t := Test{int: 1, string: "hello"}

	ptr := &t
	//fmt.Println(ptr)
	//fmt.Println(*ptr)

	ptr.int = 100
	ptr.string = "Hi"

	fmt.Println(t)
}

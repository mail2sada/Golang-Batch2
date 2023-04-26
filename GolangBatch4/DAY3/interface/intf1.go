package main

import "fmt"

type Display interface {
	Print()
}

type Read interface {
	Create()
}

type Combined interface {
	Display
	Read
}

func main() {

	fmt.Println("Nested interfaces")
}

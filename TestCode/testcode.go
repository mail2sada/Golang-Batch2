package main

import "fmt"

type Hello struct {
	Id   int
	Name string
}

func main() {
	fmt.Println("Checking how to allocate struct")

	hello := make(struct Hello)
}

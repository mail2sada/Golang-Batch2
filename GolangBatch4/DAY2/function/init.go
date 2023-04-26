package main

import "fmt"

func init() {
	fmt.Println("Calling my init function")
}

func init() {
	fmt.Println("Second init")
}

func init() {
	fmt.Println("Third init")
}

func main() {
	fmt.Println("Demo: Init....")
	fmt.Println("Inside main")
}

package main

import "fmt"

func main() {
	fmt.Println("Demo: Annonynous func")

	func(msg string) {
		fmt.Println(msg)
	}("Hello")

	fmt.Println("Exting main")
}

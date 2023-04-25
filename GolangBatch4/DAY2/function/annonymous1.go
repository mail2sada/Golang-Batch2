package main

import "fmt"

func HigherOrderFunc(slice []string, fn func(s string)) {
	for _, val := range slice {
		fn(val)
	}
}

func main() {

	fmt.Println("Annonymous functions")

	sl := []string{"Hello", "hi", "test", "best"}

	fn := func(msg string) {
		fmt.Println(msg)
	}
	HigherOrderFunc(sl, fn)

	// fn("Hello")

	// fn("Hi")

	// fn("Test")
}

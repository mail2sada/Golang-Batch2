package main

import "fmt"

// Higher order function

func IterateSlice(slice []string, fn func(val string)) {
	for _, val := range slice {
		fn(val)
	}
}

func abc(testfn func(str string)) {
	testfn("Called from abc: Hello")
}

func main() {

	fmt.Println("Annonymous functions")

	fn := func(msg string) {
		fmt.Println("In Annonymous function")
		fmt.Println(msg)
	}

	IterateSlice([]string{"Test", "123", "abc", "xyz"}, fn)

	IterateSlice([]string{"A", "B", "C", "D"}, func(msg string) {
		fmt.Println(msg)
	})
}

package main

import "fmt"

func IterateSlice(slice []string, fn func(msg string)) {
	for _, msg := range slice {
		fn(msg)
	}
}

func main() {
	fmt.Println("Demo: Higher order function")

	function := func(str string) {
		fmt.Println(str)
	}

	stringSlice := []string{"This", "is", "Mavenir", "go", "training"}
	var a = 10
	IterateSlice(stringSlice, function)

	IterateSlice(stringSlice, func(msg string) {
		a++
		msg = "[" + msg + "]"
		fmt.Println(msg)
	})

	fmt.Println(a)
}

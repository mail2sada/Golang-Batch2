package main

import "fmt"

func sliceIterator(slice []string, myfunc func(string)) {
	fmt.Println("Inside sliceIterator")

	for _, val := range slice {
		myfunc(val)
	}
}

func main() {

	fmt.Println("Demo: Ano=nonymous Functions")

	func() {
		fmt.Println("This is my anonymous funtions...")
	}()

	var myFunc = func() {
		fmt.Println("This is my annonymous functions")
	}

	stringSlice := []string{"Hi", "Hello", "Test", "Best", "123"}

	fn := func(str string) {
		fmt.Println("Received string...", str)
	}

	// sliceIterator(stringSlice, func(slice string) {
	// 	fmt.Println("Received my string...", slice)
	// })

	sliceIterator(stringSlice, fn)

	myFunc()
}

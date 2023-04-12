package main

// this code will define a function and will demonstrait how to call/ use the function
import "fmt"

// func is keyword
// MyFirstFunction is name of the function, and is identifier as well.

func MyFirstFunction() {
	fmt.Println("MyfirstFunction: is called...")
}

func main() {
	fmt.Println("Demo: Functions in Golang...")
	MyFirstFunction() // we are calling function here, the code inside function Myfirst Function will be called here
	MyFirstFunction() // we have called MyFirstFunction here... We are reusing the code inside MyFirstFunction
	// However this code is printing MyFirst function is called
	// How to make it generic and should print, what I want, istead of a fixed statement
	// We will see how to achieve this in parameter.go
}

package main

import "fmt"

func MyParameterisedFunction(str string) {
	//fmt.Println("This is my parameterised function...")
	fmt.Println(str) // this will print whatever is passed to it.
}

func main() {
	fmt.Println("Demo: Parameters to functions")
	MyParameterisedFunction("Welcome to functions")
	MyParameterisedFunction("Wow this is doing different job based on what I pass")
}

package main

import (
	"fmt"
)

func main() {
	// var variable_name type = expression
	var myInt int = 10
	// var is keyword,
	//myInt is identifier/variable
	//name and int is data type

	fmt.Println("Demo: Variable declaration")

	fmt.Println("Value of myInt is ", myInt)
	// Printing value of myInt using Println function of fmt package

	var Integer = 500

	fmt.Printf("Type of my string is %T\n", Integer)
	fmt.Println("Value of Integer is ", Integer)

	var greet = "Welcome to go training..."

	fmt.Printf("Type of greet is %T", greet)

	fmt.Println("Value of greet is ", greet)

}

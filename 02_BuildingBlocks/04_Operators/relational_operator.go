package main

import "fmt"

func main() {
	fmt.Println("Demo relational operators...")

	var a = 50

	var b = 60

	fmt.Println("Value of a is ", a, "\nValue of b is ", b)

	var result = a == b

	fmt.Println("Result of  a == b is", result)

	result = a != b

	fmt.Println("Result of  a != b is", result)

	result = a > b

	fmt.Println("Result of  a > b is", result)

	result = a < b

	fmt.Println("Result of  a < b is", result)

	result = a >= b

	fmt.Println("Result of  a >= b is", result)

	result = a <= b

	fmt.Println("Result of  a <= b is", result)

}

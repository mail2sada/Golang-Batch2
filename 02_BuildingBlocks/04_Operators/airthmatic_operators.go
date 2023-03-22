package main

import "fmt"

func main() {
	fmt.Println("Demo: Operators...")

	var a int = 500

	var b int = 100

	fmt.Println("Value of a is ", a, "\nValue of b is ", b)

	var result = a + b

	fmt.Println("Result of a, b addition is ", result)

	result = a - b

	fmt.Println("Result of a, b substraction is ", result)

	result = a * b

	fmt.Println("Result of a, b multiplication is ", result)

	result = a / b

	fmt.Println("Result of a, b division is ", result)

	result = a % b

	fmt.Println("Result of a, b remainder is ", result)

}

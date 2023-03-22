package main

import "fmt"

func main() {

	fmt.Println("Demo logical operators")

	var a = 100

	var b = 100

	var c = 200

	fmt.Println("Value of a is ", a, "\n Value of b is ", b, "\nValue of c is ", c)

	var result = (a == b) && (a == c)

	// here first condition is true and second condition is false, and operation should result in false
	fmt.Println("Result of (a == b) && (a == c) ", result)

	result = (a == b) || (a == c)

	// here first condition is true and second condition is false, or operation should result in true
	fmt.Println("Result of (a == b) || (a == c) ", result)

	result = !((a == b) && (a == c))

	// here first condition is true and second condition is false, and operation should result in false, however a not operation on whole expresion should result in true
	fmt.Println("Result of !((a == b) && (a == c)) ", result)

}

package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Demo: Operators")

	var v1, v2 int

	fmt.Println("Enter value of v1")
	fmt.Scanf("%d", &v1)

	fmt.Println("Enter value for v2")
	fmt.Scanf("%d", &v2)

	// airthmatic operators
	var sum = v1 + v2

	var diff = v1 - v2

	var product = v1 * v2

	var divident = v1 / v2

	var remainder = v1 % v2

	fmt.Println("Sum of v1 and v2 is ", sum, "\nDifference of v1 and v2 is ", diff, "\nProduct of V1 and Ve is ", product)

	fmt.Println("Divident of v1 and v2 ", divident)

	fmt.Println("Remainder of v1 and v2", remainder)

	// relational operators

	var equal = v1 == v2

	var greater = v1 > v2

	var less = v1 < v2

	var notEqual = v1 != v2

	fmt.Println("Type of Equal is ", reflect.TypeOf(equal))

	fmt.Println(equal, greater, less, notEqual)

}

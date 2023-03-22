package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Short declaration operator...")

	a := 10

	fmt.Println("Type of a is ", reflect.TypeOf(a))

	fmt.Println("Value of a is ", a)

	// we are declaring to variable here

	x, y := 100, "Welcome to go training!!!"

	fmt.Println("Type of x is ", reflect.TypeOf(x))
	fmt.Println("value of x is ", x)

	fmt.Println("Type of y is ", reflect.TypeOf(y))
	fmt.Println("Value of y is ", y)

	// previously delcared x will be initialised here..
	x, z := 100, 5.5

	fmt.Println("Type of z is ", reflect.TypeOf(z))

	fmt.Println("Value of z is ", z)

}

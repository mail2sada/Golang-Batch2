package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Demo: Variable declaration")

	var variable_1 int = 100

	fmt.Printf("\n Value of %d", variable_1)

	fmt.Printf("\n Type of variable_1 is %T\n", variable_1)

	var variable_2 = "Hello world..."

	fmt.Println("Value of variable_2", variable_2)

	fmt.Printf("Value of variable_2 is %s", variable_2)
	fmt.Printf("Type of variable_2 is %T\n", variable_2)

	var v1, v2, v3, v4 uint = 10, 20, 30, 40

	fmt.Println(reflect.TypeOf(v1), reflect.TypeOf(v2))

	fmt.Println(v1, v2, v3, v4)

	var a1, a2, a3, a4 = 10, 20, 30, 40

	fmt.Println(a1, a2, a3, a4)

	var integer, strings, float = 10, "Hello", 10.55

	fmt.Println(integer, strings, float)

	fmt.Println(reflect.TypeOf(integer), reflect.TypeOf(strings), reflect.TypeOf(float))

	const const1 int = 10

	const const2, const3 = 10, "Test"

	fmt.Println(const1, const2, const3)

	fmt.Println(reflect.TypeOf(const1), reflect.TypeOf(const2), reflect.TypeOf(const3))

}

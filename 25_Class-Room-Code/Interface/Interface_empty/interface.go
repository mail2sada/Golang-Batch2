package main

import (
	"fmt"
	"reflect"
)

func MyPrint(elements ...interface{}) {
	//fmt.Println(elements...)

	for _, val := range elements {
		fmt.Println(reflect.TypeOf(val))
		fmt.Print(val)
	}
}

func add(a, b interface{}) interface{} {

	switch a.(type) {
	case int:
		c := a.(int) + b.(int)
		return c
	case string:
		c := a.(string) + b.(string)
		return c
	case float64:
		c := a.(float64) + b.(float64)
		return c
	default:
		c := -100
		return c
	}

}

func main() {
	fmt.Println("Demo interface...")
	MyPrint(10, 10.25, "Hello")

	MyPrint(add(10, 20))

	MyPrint(add(10.1, 10.25))

	MyPrint(add("Hello", "Mavenir.."))

	//MyPrint(add("Hello", 10))
}

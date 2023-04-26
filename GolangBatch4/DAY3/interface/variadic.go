package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Render(args ...interface{}) {
	fmt.Println(args...)

	for _, val := range args {
		fmt.Println(val)
	}
}

func Add(a interface{}, b interface{}) (res interface{}, err error) {

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		err = errors.New("a and b are of different types")
		return
	}
	switch a.(type) {
	case int:
		res = a.(int) + b.(int)
	case float64:
		res = a.(float64) + b.(float64)
	case float32:
		res = a.(float32) + b.(float32)
	case string:
		res = a.(string) + b.(string)

	default:
		err = errors.New("type not handled")
		return

	}
	return
}

func main() {
	fmt.Println("Variadic with Interface")

	Render(10, "Hello", 3.142, true)

	a, _ := Add(10, 20)
	fmt.Println(a)

	b, err := Add("Hello", 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}

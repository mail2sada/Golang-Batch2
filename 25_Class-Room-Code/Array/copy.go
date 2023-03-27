package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Array copy")

	var array1 = [3]string{"Hello", "how", "are you?"}

	array2 := array1

	array2[1] = "abcdef"

	fmt.Println(array1)
	fmt.Println(array2)

	array3 := &array1

	fmt.Println(reflect.TypeOf(array3))

	array3 = &array2

	array3[0] = "Test"

	fmt.Println(array1)
	fmt.Println(array3)
}

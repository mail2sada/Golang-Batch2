package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Demo: Strings package")

	var a = "Hello"
	var b = "Hello"

	if a == b {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	if reflect.DeepEqual(a, b) {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	if strings.Compare(a, b) == 0 {
		fmt.Println("a and b are equal..")
	} else {
		fmt.Println("a and b are not equal")
	}

}

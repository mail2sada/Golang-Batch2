package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	str1 := "Hello"
	str2 := "hello"
	// != cam also be used
	if str1 == str2 {
		fmt.Println("Equal")
	} else if str2 < str1 {
		fmt.Println("str1 less than str2")
	} else {
		fmt.Println("str2 less than str1")
	}

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("Equal")
	}

	fmt.Println(strings.Compare(str1, str2))

	if reflect.DeepEqual(str1, str2) {
		fmt.Println("String are equal..")
	} else {
		fmt.Println("Strings are not equal")
	}

	if strings.EqualFold(str1, str2) {
		fmt.Println("Strings are equal by ignoring case...")
	} else {
		fmt.Println("Not equal...")
	}

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Demo strconv")
	a := 100
	//const pi = 3.142

	astr := strconv.Itoa(a)

	str := "Value of a is " + astr

	fmt.Println(str)

	//"Value of a is 100"
}

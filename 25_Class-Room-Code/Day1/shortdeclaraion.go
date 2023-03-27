package main

import "fmt"

var (
	global  int    = 100
	string1 string = ""
	float   float32
)

func main() {

	// var a = 100

	a := 100

	integer, strings := 10, "hello"

	fmt.Println(a, integer, strings)
}

package main

import "fmt"

//& - refrecing
//* Dereferencing

func main() {

	fmt.Println("Demo: Pointers")

	var a = 100

	var ptr *string 

	fmt.Println(ptr)

	fmt.Println(*ptr)

	*ptr = 500

	fmt.Println(*ptr)
	fmt.Println(a)
}

package main

import "fmt"

func main() {
	fmt.Println("Demo: Pointers")
	var a int = 100

	var ptr *int = &a

	var ptrToPtr **int = &ptr

	fmt.Println(ptr)
	fmt.Println(&ptr)
	fmt.Println(*ptr)
	fmt.Println(a)
	fmt.Println(ptrToPtr)   //line14
	fmt.Println(*ptrToPtr)  //line13
	fmt.Println(**ptrToPtr) //line 15 16
}

package main

import "fmt"

func main() {
	fmt.Println("Demo: Arrays....")

	// integer array that can hold 3 elements

	var intArray [3]int = [3]int{}

	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3

	fmt.Println(intArray)

}

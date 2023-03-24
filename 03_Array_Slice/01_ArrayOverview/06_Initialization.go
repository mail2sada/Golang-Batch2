package main

import "fmt"

func main() {
	fmt.Println("Demo: Initializing specific elements of array")

	// we will declare an array of int with size 10 and intializing index 0, 5, 9

	var integerArray = [10]int{0: 100, 5: 200, 9: 300}

	fmt.Println("Contents of array are ", integerArray)
}

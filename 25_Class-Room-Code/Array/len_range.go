package main

import "fmt"

func main() {
	fmt.Println("Demo: Arrays")

	var testArray = [5]int{1, 2, 3, 4, 5}

	fmt.Println(len(testArray))

	for i := 0; i < len(testArray); i++ {
		fmt.Println("Value of testArray ", testArray[i])
	}

	for i, v := range testArray {
		fmt.Println("Value of testArray ", i, v)
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Demo: Summry of Arrays...")

	// Declaring arrays

	var array [3]int = [3]int{1, 2, 3}

	fmt.Println("Contents of array:", array)

	fmt.Println("Updating element at index 1")
	array[1] = 100.

	fmt.Println("Contents of array after update:", array)

	var myStringArray = [3]string{"Hello", "Good", "Morning"}

	fmt.Println("Contents of myStringArray:", myStringArray)

	arrayInt := [3]int{6, 3, 9}

	fmt.Println("Contents of arrayInt", arrayInt)

	fmt.Println("Iterating array using len function")

	for i := 0; i < len(myStringArray); i++ {
		fmt.Println("Contents of myStringArray at index", i, " is ", myStringArray[i])
	}

	fmt.Println("Iterating array using range operator")

	for idx, val := range myStringArray {
		fmt.Println("Contents of myStringArray at index:", idx, " is:", val)
	}
}

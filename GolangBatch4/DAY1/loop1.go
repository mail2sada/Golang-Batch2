package main

import "fmt"

func main() {
	fmt.Println("Demo: loop")

	var myString = "Welcome to Mavenir Go lang training"

	// for idx, val := range myString {

	// 	fmt.Printf("\n Value at %d is %c", idx, val)
	// 	// fmt.Println("\nValue at ", idx, "is", val)
	// 	// fmt.Printf("\n%c", val)
	// }

	for idx, _ := range myString {

		fmt.Printf("\n Value at %d", idx)
		// fmt.Println("\nValue at ", idx, "is", val)
		// fmt.Printf("\n%c", val)
	}
}

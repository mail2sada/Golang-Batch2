package main

import "fmt"

func main() {

	fmt.Println("Demo: if-else")

	myInt := 100

	testInt := 50

	if myInt > testInt {
		fmt.Println("myInt is greater than testInt")
	} else {
		fmt.Println("myInt is not greater than testInt")
	}

	if myInt < testInt {
		fmt.Println("myInt is less than testInt")
	} else {
		fmt.Println("myInt is not less than testInt")
	}
}

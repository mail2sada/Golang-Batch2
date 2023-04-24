package main

import "fmt"

func main() {
	fmt.Println("Demo: String array")

	var stringArray [3]string = [3]string{"Welcome", "to", "Mavenir"}

	fmt.Println(stringArray)

	for idx, val := range stringArray {
		fmt.Println(idx, val)
	}

	stringArray[1] = "TO"
	fmt.Println(stringArray)
}

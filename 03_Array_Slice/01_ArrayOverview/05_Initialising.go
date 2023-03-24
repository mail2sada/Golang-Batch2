package main

import "fmt"

func main() {

	fmt.Println("Demo: Array Initialisation 1...")

	var intArray [3]int = [3]int{1, 2, 3}

	fmt.Println("Contents of intArray are:", intArray)

	var strArray = [3]string{"Hello ", "Hi ", "How are you? "}

	fmt.Println("Contents of str Array", strArray)

	uintArray := [3]uint8{10, 20, 30}

	fmt.Println("Contents of uintArray", uintArray)
}

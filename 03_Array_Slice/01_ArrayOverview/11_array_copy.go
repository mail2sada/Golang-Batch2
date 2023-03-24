package main

import "fmt"

func main() {

	fmt.Println("Demo copying array...")

	var array [5]string = [5]string{"Golang", "is", "a", "wonderful", "language.."}

	var copyArray = array

	fmt.Println("Contents of array are:", array)

	fmt.Println("Contents of copyArray are:", copyArray)

	copyArray[3] = "I have changed this..."

	fmt.Println("Printing array and copyArray after changing copyArray[3]=\"I have changed this...\"")

	fmt.Println("Contents of array are:", array)

	fmt.Println("Contents of copyArray:", copyArray)
}

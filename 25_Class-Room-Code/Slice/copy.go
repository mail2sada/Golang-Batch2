package main

import "fmt"

func main() {
	fmt.Println("Demo: Copy slice...")

	sliceString := []string{"Hello", "12345", "Mavenir India..."}

	/*copySlice := sliceString

	copySlice[1] = "Go lang training..."

	fmt.Println(copySlice)

	fmt.Println(sliceString)
	*/

	var copySlice []string = make([]string, 5, 5)

	copy(copySlice, sliceString)

	copySlice[1] = "Go training..."

	fmt.Println(copySlice)
	fmt.Println(sliceString)

}

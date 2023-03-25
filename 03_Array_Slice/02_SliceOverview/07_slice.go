package main

import (
	"fmt"
)

func main() {

	fmt.Println("Demo: Delete item in slice")

	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Contents of slice are...", sliceInt)

	// we are deleting item at index idx

	idx := 5
	// item at index 5 is 6
	sliceInt = append(sliceInt[:idx], sliceInt[idx+1:]...)

	fmt.Println("After deleting element at index 6, contents of slice are:", sliceInt)

}

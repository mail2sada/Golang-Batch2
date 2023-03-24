package main

import "fmt"

func main() {
	fmt.Println("Demo: Multi-dimentsional array")

	var multiDimensionalArray [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %d ", multiDimensionalArray[i][j])
		}
		fmt.Println()
	}

}

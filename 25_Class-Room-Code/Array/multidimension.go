package main

import "fmt"

func main() {
	fmt.Println("Demo: Multi dimension")

	var multiArray [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for _, val := range multiArray {
		for _, val1 := range val {
			fmt.Println(val1)
		}
	}
}

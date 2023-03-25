package main

import "fmt"

func main() {
	fmt.Println("Multi Dimentsional Arrays...")

	var array [3][5]int = [3][5]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for _, val := range array {

		for _, val1 := range val {
			//fmt.Println("Column number", iidx)
			fmt.Print(val1)
		}
		println()
	}

}

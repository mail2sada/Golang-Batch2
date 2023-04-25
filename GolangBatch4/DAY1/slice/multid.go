package main

import "fmt"

func main() {
	fmt.Println("Demo Multi Dimensional Array")

	var myslice [][]int = [][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9}, {10, 11, 12, 13, 14, 15, 16}}

	for idx, sublice := range myslice {
		fmt.Println(idx, sublice)
		for id, val := range sublice {
			fmt.Println("loc[]", idx, id, "]:", val)
		}
	}
}

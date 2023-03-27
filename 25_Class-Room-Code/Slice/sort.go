package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Demo: sort")

	var in int

	slice := []int{4, 2, 6, 1, 8, 9, 3, 5, 7, 100: 10}

	fmt.Println(slice)

	res := sort.SearchInts(slice, 6)

	fmt.Println(res)

	if sort.IntsAreSorted(slice) {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted...")
	}
	sort.Ints(slice)

	if sort.IntsAreSorted(slice) {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted...")
	}
	fmt.Println(slice, "after sort")

	res = sort.SearchInts(slice, 6)

	fmt.Println(res)

}

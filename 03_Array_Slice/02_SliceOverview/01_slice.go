package main

import "fmt"

func main() {

	fmt.Println("Demo: Slice Declaration...")

	var intSlice []int = []int{1, 2, 3}

	intSlice = append(intSlice, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Length of the slice:", len(intSlice))
	fmt.Println("Capacity of the slice:", cap(intSlice))

	fmt.Println("Accessing element at index [3]:", intSlice[3])

	fmt.Println("Iterating slice using len")
	for i := 0; i < len(intSlice); i++ {
		fmt.Println("Content at index [", i, "]  = ", intSlice[i])
	}

	fmt.Println("Iterating slice using range operator")
	for idx, val := range intSlice {
		fmt.Println("Content at index [", idx, "]  = ", val)

	}

}

package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice delete...")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	idx := 3

	slice = append(slice[:idx], slice[idx+1:]...)

	fmt.Println(slice, "After delete..")

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for idx, val := range slice {
		fmt.Println(idx, val)
	}
}

// make, len, cap, append,

// updating, deleting

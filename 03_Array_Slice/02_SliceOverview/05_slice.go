package main

import "fmt"

func main() {
	fmt.Println("Demo: Usage of make in slice...")

	// statement below will allocate a string slice of lenght 5
	var stringSlice = make([]string, 5)

	// here both capacity and length should be same

	fmt.Println("Length of stringSlice:", len(stringSlice))
	fmt.Println("Capacity of stringSlice:", cap(stringSlice))

	fmt.Println("\nWe are re allocating stringSlice using length and capacity")

	stringSlice = make([]string, 5, 10)
	fmt.Println("Length of stringSlice:", len(stringSlice))
	fmt.Println("Capacity of stringSlice:", cap(stringSlice))
}

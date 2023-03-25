package main

import "fmt"

func main() {
	fmt.Println("Demo: functions to handle slice, make, len, cap and append")

	// using make to create a slice with length 5

	stringSlice := make([]string, 0, 5)

	// let's print slice capacity and length.

	fmt.Println("Capacity of stringSlice:", cap(stringSlice))

	fmt.Println("length of stringSlice:", len(stringSlice))

	// using make to create a slice with length and capacity

	newSlice := make([]string, 0, 10)

	// Reading length of the slice,

	fmt.Println("Length of newSlice is:", len(newSlice))

	// Reading capacity of the slice

	fmt.Println("Capacity of newSlice is ", cap(newSlice))

	// appending items to the slice

	stringSlice = append(stringSlice, "A", "B", "C", "D", "E")

	newSlice = append(newSlice, "F", "G", "H", "I", "J")

	fmt.Println("Contents of stringSlice:", stringSlice)
	for _, val := range stringSlice {
		fmt.Println(val)
	}

	fmt.Println("Contents of newSlice:", newSlice)

	// concatinating 2 slices

	stringSlice = append(stringSlice, newSlice...)

	fmt.Println("Contets of stringSlice after concatinating newSlice", stringSlice)

	// concatingting a slice and array

	array := [5]string{"K", "L", "M", "N", "O"}

	stringSlice = append(stringSlice, array[:]...)

	fmt.Println("Contents of string slice after appending a array", stringSlice)

	fmt.Println("Capacity of stringSlice", cap(stringSlice))

	fmt.Println("Length of stringSlice", len(stringSlice))

}

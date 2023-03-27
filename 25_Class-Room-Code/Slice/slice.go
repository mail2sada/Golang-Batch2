package main

import "fmt"

func main() {
	fmt.Println("Demo slice")

	var slice []int = []int{1, 2, 3}

	fmt.Println("Length of slice is ", len(slice))

	fmt.Println("Capacity of slice ", cap(slice))

	var mySlice []int = make([]int, 100, 100)

	mySlice[0] = 500

	fmt.Println("Length of mySlice ", len(mySlice))

	fmt.Println("Capacity of mySlice is", cap(mySlice))

	mySlice = append(mySlice, 100)

	fmt.Println("Length of mySlice ", len(mySlice))

	fmt.Println("Capacity of mySlice is", cap(mySlice))

	fmt.Println(mySlice)

}

// Declaring
// Appending
// Updating
// Deleting

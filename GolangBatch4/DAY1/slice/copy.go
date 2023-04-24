package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice copy")

	var mySlice = []int{1, 2, 3}
	fmt.Println("Cap", cap(mySlice))
	var myArray = [3]int{}
	myArray = [3]int(mySlice[:3])
	fmt.Println(myArray)

	copySlice := mySlice

	fmt.Println("Copyslice", copySlice)
	fmt.Println("OriginalSlice", mySlice)

	copySlice[0] = 100
	fmt.Println("After copy")
	fmt.Println("Copyslice", copySlice)
	fmt.Println("OriginalSlice", mySlice)
	mySlice = append(mySlice, 10, 11, 12)
	fmt.Println("After append")

	fmt.Println("Copyslice", copySlice)
	fmt.Println("OriginalSlice", mySlice)
	copySlice[1] = 10000

	fmt.Println("Copyslice", copySlice)
	fmt.Println("OriginalSlice", mySlice)
}

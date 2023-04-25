package main

import "fmt"

func main() {
	fmt.Println("Demo: Insert slice")

	var intSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intSlice)
	index := 5

	intSlice = append(intSlice[:index+1], intSlice[index:]...)
	intSlice[index] = 100
	fmt.Println(intSlice)
}

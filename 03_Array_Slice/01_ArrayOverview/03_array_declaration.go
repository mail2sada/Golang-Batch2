package main

import "fmt"

func main() {
	fmt.Println("Demo: Array declaration")
	array := [3]int{}

	array[0] = 100
	array[1] = 200
	array[2] = 300

	fmt.Println("Content of array are:", array)
}

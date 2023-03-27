package main

import "fmt"

func main() {
	fmt.Println("Demo: filterning")

	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array[:5])
	fmt.Println(array[5:])

	fmt.Println(array[3:7])
}

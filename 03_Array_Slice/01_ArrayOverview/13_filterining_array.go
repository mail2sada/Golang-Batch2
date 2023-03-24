package main

import "fmt"

func main() {

	fmt.Println("Demo: Filtering elements of array..")

	var array [10]int = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println("Filering out first 5 elements", array[:5])

	fmt.Println("Fileterning out last 5 elements", array[5:])
	fmt.Println("Filterning out from index 3 to 7", array[3:8])
}

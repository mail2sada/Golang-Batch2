package main

import "fmt"

func main() {

	fmt.Println("Demo: len function on collection")

	var collectionInt [10]int = [10]int{1, 2, 3, 4, 5}

	fmt.Println("Length of collectionInt is:", len(collectionInt))
}

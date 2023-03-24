package main

import "fmt"

func main() {

	fmt.Println("Demo: Iterating over a collection")

	var collectionInt = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(collectionInt); i++ {
		fmt.Printf("collectionInt[%d]=%d\n", i, collectionInt[i])
	}
}

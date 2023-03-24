package main

import "fmt"

func main() {
	fmt.Println("Demo: iterating loop using range")

	var strArray = [10]string{"Good", "Morning", "How", "Are", "You", "Wonderful", 9: "Amazing"}

	for idx, val := range strArray {
		fmt.Println("strArray @Index:", idx, " is:", val)
	}
}

package main

import "fmt"

func main() {

	fmt.Println("Demo: for loop with condition")

	i := 10

	for i >= 0 {
		fmt.Println("Value of i is ", i)
		i--
	}
}

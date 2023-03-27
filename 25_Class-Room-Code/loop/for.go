package main

import "fmt"

func main() {

	fmt.Println("Demo: For loop")

	for i := 0; i <= 100; i++ {
		fmt.Println("Value of i", i)
	}

	for i := 0; i <= 100; i += 2 {
		fmt.Println(i)
	}

	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}
}

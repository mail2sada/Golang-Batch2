package main

import "fmt"

func main() {

	fmt.Println("Demo: Loop control statement continue")

	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			continue
		}
		fmt.Println("Value of i:", i)
	}

	i := 0

	for i < 100 {
		i++
		if i%2 != 0 {
			continue
		}
		fmt.Println("Value of i:", i)
	}
}

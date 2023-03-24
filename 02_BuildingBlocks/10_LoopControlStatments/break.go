package main

import "fmt"

func main() {
	fmt.Println("Demo: loop control statements break")

	i := 0

	for {
		i++

		if i == 50 {
			break
		}
		fmt.Println("Value of i:", i)
	}
}

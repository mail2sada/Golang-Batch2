package main

import "fmt"

func main() {
	fmt.Println("Demo: loop control statement goto")

Label2:
	fmt.Println("Hello...")
	x := 0

Label1:
	fmt.Println("Starting loop")
	for x < 10 {
		x++
		fmt.Println("Value of x:", x)
		if x > 5 {
			fmt.Println("Going to label 1")
			goto Label1
		}

		if x > 8 {
			fmt.Println("Goining to label2")
			goto Label2
		}

	}
}

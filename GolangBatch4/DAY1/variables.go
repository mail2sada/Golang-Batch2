package main

import "fmt"

func main() {
	var variable1 int = 100

	fmt.Println("Value of variable1 is", variable1)
	variable1 = 500

	fmt.Println("New value os variable1 is", variable1)

	var variable2 = 500

	fmt.Printf("\nType variable2 is %T", variable2)

	var variable3 = 3.142
	fmt.Printf("Type variable3 is %T", variable3)

	var a, b, c int = 10, 20, 30

	fmt.Printf("\n%d %d %d", a, b, c)

	var x, y, z = 10, 3.142, "Hello"

	fmt.Println("\n", x, y, z)

}

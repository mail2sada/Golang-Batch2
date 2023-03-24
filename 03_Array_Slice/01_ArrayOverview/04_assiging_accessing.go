package main

import "fmt"

func main() {
	fmt.Println("Demo: Assigining and accessing arrays")
	var color [5]string = [5]string{}

	// Assigining array elements
	color[0] = "Red"
	color[1] = "Green"
	color[2] = "Blue"
	color[3] = "Orrange"
	color[4] = "Purple"

	fmt.Println("Contents of color array:", color)

	// Accessing array elements...

	fmt.Println("color[0]:", color[0])
	fmt.Println("color[1]:", color[1])
	fmt.Println("color[2]:", color[2])
	fmt.Println("color[3]:", color[3])
	fmt.Println("color[4]:", color[4])

}

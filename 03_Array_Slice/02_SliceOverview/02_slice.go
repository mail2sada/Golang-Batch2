package main

import "fmt"

func main() {

	fmt.Println("Demo: Declaration of slice using make...")

	var myStringSlice []string = make([]string, 5)

	fmt.Println("Length of myStringSlice:", len(myStringSlice))
	fmt.Println("Capacity of myStringSlice:", cap(myStringSlice))

	//populating list with few elements.

	myStringSlice = append(myStringSlice, "Hello", "how", "are", "you")

	fmt.Println("Length of myStringSlice:", len(myStringSlice))
	fmt.Println("Capacity of myStringSlice:", cap(myStringSlice))

	fmt.Println("Iterating through the slice using len")

	for i := 0; i < len(myStringSlice); i++ {

		fmt.Println("Contents of myStringSlice[", i, " ]", myStringSlice[i])
	}

	fmt.Println("Iterating through the slice using range")

	for idx, val := range myStringSlice {
		fmt.Println("Contents of myStringSlice[", idx, " ]", val)

	}

}

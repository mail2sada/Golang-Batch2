package main

import "fmt"

func main() {

	fmt.Println("Demo Map iteration")

	var emp map[string]int = map[string]int{"Anil": 30000, "Vijay": 50000, "Santhosh": 45000, "Manoj": 45000}

	for key, val := range emp {
		fmt.Println("Key is", key, "Value is ", val)
	}
}

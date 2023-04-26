package main

import "fmt"

func main() {
	fmt.Println("Demo: defer")
	// OPened few files
	// network connection
	defer func() {
		// Close all files
		// Close all connection
		// all house keeping
		fmt.Println("This is annonynous function called with defer")
	}()

	fmt.Println("Exit")
}

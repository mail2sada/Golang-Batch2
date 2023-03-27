package main

import "fmt"

func main() {
	fmt.Println("For with range..")

	var welcome = "Welcome to Go training"

	for index, val := range welcome {
		fmt.Printf("Welcome[%d]=%c\n", index, val)
	}

	for _, value := range welcome {
		fmt.Println(value)
	}

}

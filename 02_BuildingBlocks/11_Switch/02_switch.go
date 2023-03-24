package main

import "fmt"

func main() {
	fmt.Println("Demo: Demo Switch String...")

	testString := "Hello"
	switch testString {
	case "Hello":
		fmt.Println("Its english")
	case "Namaste":
		fmt.Println("Its Hindi")
	default:
		fmt.Println("I am yet to learn new language..")
	}
}

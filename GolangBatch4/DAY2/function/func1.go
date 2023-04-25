package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to Go training")
}

func printMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Demo: Functions")

	printWelcomeMessage()

	printWelcomeMessage()

	printWelcomeMessage()

	printMessage("Hello this is functions demo...")

	printMessage("We are testing functions with argument")

	printMessage("This is go functions session")
}

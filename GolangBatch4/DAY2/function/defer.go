package main

import "fmt"

func TestDefer() {
	fmt.Println("Testing defer")

	defer fmt.Println("Deferining this print in TestDefer..")

	fmt.Println("Exiting TestDefer")
}

func main() {
	fmt.Println("Demo defer...")
	defer TestDefer()
	defer fmt.Println("This print is defered")

	defer fmt.Println("My second defer")

	fmt.Println("We are in main")

	fmt.Println("Exiting main")

}

package main

import "fmt"

func main() {

	fmt.Println("Demo: Constant declaration")

	const pi float64 = 3.14159

	fmt.Printf("Type of pi is %T", pi)

	const milliSecond = 1
	const second = 1000 * milliSecond
	const minute = 60 * second
	const hour = 60 * minute

	const greeting = "Welcome to go training..."

	fmt.Println("Value of milliSecond:", milliSecond)
	fmt.Println("Value of second:", second)
	fmt.Println("Value of minute:", minute)
	fmt.Println("Value of hour:", hour)

	fmt.Println(greeting)
}

package main

import "fmt"

func main() {

	var dayOfMonth int = 0

	fmt.Println("Enter day of Month")
	fmt.Scanf("%d", &dayOfMonth)
	switch {
	case dayOfMonth > 0 && dayOfMonth <= 7:
		fmt.Println("First week of Month")
	case dayOfMonth > 7 && dayOfMonth <= 14:
		fmt.Println("Second Week")
	case dayOfMonth > 15 && dayOfMonth <= 22:
		fmt.Println("third Week")
	case dayOfMonth > 22 && dayOfMonth <= 31:
		fmt.Println("last Week")
	default:
		fmt.Println("Invalid input")
	}
}

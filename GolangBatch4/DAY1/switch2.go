package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch")

	var weekDay string
	fmt.Println("Please enter weekday")
	fmt.Scanf("%s", &weekDay)
	switch weekDay {
	case "MON", "mon":
		fmt.Println(1)
	case "TUE", "tue":
		fmt.Println(2)

	case "WED":
		fmt.Println(3)
	default:
		fmt.Println("Full case statement not implemented")

	}
}

package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch")
	var weekday int

	fmt.Println("Pls enter value for weekday")
	fmt.Scanf("%d", &weekday)
	switch weekday {

	case 1:
		fmt.Println("MON")
		fallthrough
	case 2:
		fmt.Println("TUE")

	case 3:
		fmt.Println("WED")

	case 4:
		fmt.Println("THU")

	case 5:
		fmt.Println("FRI")
	case 0:
		//fmt.Println("Sun")
		fallthrough
	case 6:
		fmt.Println("Weekend")
		//fmt.Println("SAT")
	default:
		fmt.Println("Invalid input")

	}
}

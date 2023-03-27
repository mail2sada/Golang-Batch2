package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch")

	var AvgMarks int
	fmt.Println("Enter average marks")
	fmt.Scanf("%d", &AvgMarks)

	switch {
	case AvgMarks > 70:
		fmt.Println("Distinction")
	case AvgMarks > 60:
		fmt.Println("First Class")
	case AvgMarks > 50:
		fmt.Println("Second Class")
	default:
		fmt.Println("Better luck next time..")
	}
}

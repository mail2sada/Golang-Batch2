package main

import "fmt"

/*
0 Sun
1 Mon
2 Tue
3 Wed
4 Thur
5 Fri
6 Sat
*/
func main() {

	var weekDay int
	fmt.Println("Please enter value for weeday")
	fmt.Scanf("%d", &weekDay)
	// if weekDay == 0 {
	// 	fmt.Println("Sun")
	// } else {
	// 	if weekDay == 1 {
	// 		fmt.Println("Mon")
	// 	}
	// }

	if weekDay == 0 {
		fmt.Println("Sun")
	} else if weekDay == 1 {
		fmt.Println("Mon")
	} else if weekDay == 2 {
		fmt.Println("Tue")
	} else if weekDay == 3 {
		fmt.Println("Wed")

	} else if weekDay == 4 {
		fmt.Println("Thu")

	} else if weekDay == 5 {
		fmt.Println("Fri")
	} else if weekDay == 6 {
		fmt.Println("Sat")

	} else {
		fmt.Println("Invalid input")
	}

}

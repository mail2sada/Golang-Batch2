package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Demo: if")

	var myInt int

	fmt.Println(myInt)

	fmt.Println("Enter value for myInt")
	fmt.Scanf("%d", &myInt)

	if myInt > 0 {
		fmt.Println("myInt is + ve ")
	} else {
		fmt.Println("myInt may not be +ve")
	}

	if myInt > 0 {
		fmt.Println("Number is +ve")
	} else if myInt < 0 {
		fmt.Println("Number is -ve")
	} else {
		fmt.Println("Its zero")
	}

	if day := time.Now().Day(); day > 25 {
		fmt.Println("Month end..", day)
	}

}
